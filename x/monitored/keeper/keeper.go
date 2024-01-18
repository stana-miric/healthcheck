package keeper

import (
	"fmt"
	"time"

	commonTypes "healthcheck/x/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	"github.com/tendermint/tendermint/libs/log"

	"healthcheck/x/monitored/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		channelKeeper types.ChannelKeeper
		portKeeper    types.PortKeeper
		scopedKeeper  exported.ScopedKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper exported.ScopedKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

// SetHealthcheckChannel sets the the healthcheck channel
func (k Keeper) GetHealthcheckChannel(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	channelIdBytes := store.Get(types.HealthcheckChannelKey)
	if len(channelIdBytes) == 0 {
		return "", false
	}
	return string(channelIdBytes), true
}

// SetHealthcheckChannel gets the the healthcheck channel
func (k Keeper) SetHealthcheckChannel(ctx sdk.Context, channelID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.HealthcheckChannelKey, []byte(channelID))
}

// GetLastCheckin gets the the last checkin info
func (k Keeper) GetLastCheckin(ctx sdk.Context) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastCheckinKey)
	if bz == nil {
		return 0, false
	}
	return sdk.BigEndianToUint64(bz), false
}

// SetLastCheckin sets the the last checkin info
func (k Keeper) SetLastCheckin(ctx sdk.Context, lastCheckin uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(lastCheckin)
	store.Set(types.LastCheckinKey, bz)
}

// ----------------------------------------------------------------------------
// IBC Keeper Logic
// ----------------------------------------------------------------------------

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the IBC app module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the IBC app module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the IBC app module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SendIBCPacket(ctx sdk.Context, packetData commonTypes.HealthcheckUpdateData) bool {
	channelID, ok := k.GetHealthcheckChannel(ctx)
	if !ok {
		return false
	}

	portId := k.GetPort(ctx)
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portId, channelID))
	if !ok {
		return false
	}

	packet, err := packetData.Marshal()
	if err != nil {
		return false
	}

	k.channelKeeper.SendPacket(ctx,
		channelCap,
		portId,
		channelID,
		clienttypes.Height{}, //  timeout height disabled
		uint64(ctx.BlockTime().Add(time.Hour).UnixNano()), // timeout timestamp
		packet,
	)

	return true
}
