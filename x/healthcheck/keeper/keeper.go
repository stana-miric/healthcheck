package keeper

import (
	"fmt"

	commonTypes "healthcheck/x/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	ibctm "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/libs/log"

	"healthcheck/x/healthcheck/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		channelKeeper    types.ChannelKeeper
		portKeeper       types.PortKeeper
		connectionKeeper types.ConnectionKeeper
		clientKeeper     types.ClientKeeper
		scopedKeeper     exported.ScopedKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	connectionKeeper types.ConnectionKeeper,
	clientKeeper types.ClientKeeper,
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

		channelKeeper:    channelKeeper,
		portKeeper:       portKeeper,
		connectionKeeper: connectionKeeper,
		clientKeeper:     clientKeeper,
		scopedKeeper:     scopedKeeper,
	}
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

func (k Keeper) GetClientChainIdFromChannel(ctx sdk.Context, channelID string) (string, error) {
	channel, ok := k.channelKeeper.GetChannel(ctx, k.GetPort(ctx), channelID)
	if !ok {
		return "", sdkerrors.Wrapf(types.ErrChannelNotFound, "channel not found for channel ID: %s", channelID)
	}

	return k.GetClientChainIdFromConnection(ctx, channel.ConnectionHops[0])

}

func (k Keeper) GetClientChainIdFromConnection(ctx sdk.Context, connectionID string) (string, error) {
	connection, found := k.connectionKeeper.GetConnection(ctx, connectionID)
	if !found {
		return "", sdkerrors.Wrapf(types.ErrInvalidConnection, "connection-id: %s", connectionID)
	}

	clientState, found := k.clientKeeper.GetClientState(ctx, connection.ClientId)
	if !found {
		return "", sdkerrors.Wrapf(types.ErrInvalidClient, "client-id: %s", connection.ClientId)
	}

	chainID := clientState.(*ibctm.ClientState).ChainId
	return chainID, nil
}

func (k Keeper) OnRecvHealthcheckPacket(ctx sdk.Context, channelID string, packetData commonTypes.HealthcheckPacketData) exported.Acknowledgement {
	chainID, err := k.GetClientChainIdFromChannel(ctx, channelID)
	if err != nil {
		panic(fmt.Errorf("cannot get client id from chainnel: %s", chainID))
	}

	chain, found := k.GetMonitoredChain(ctx, chainID)
	if !found {
		panic(fmt.Errorf("monitored chain not registered: %s", chainID))
	}

	chain.Status.Block = packetData.GetHealtcheckUpdate().Block
	chain.Status.Timestamp = packetData.GetHealtcheckUpdate().Timestamp
	chain.Status.RegistryBlockHeight = uint64(ctx.BlockHeader().Height)
	chain.Status.Status = string(types.Active)

	k.SetMonitoredChain(ctx, chain)

	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
	return ack
}

func (k Keeper) CloseChannel(ctx sdk.Context, channelID string) {
	portID := k.GetPort(ctx)
	channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)
	if found && channel.State != channeltypes.CLOSED {
		capName := host.ChannelCapabilityPath(portID, channelID)
		chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
		if !ok {
			k.Logger(ctx).Error("could not retrieve channel capability at: %s", capName)
		}

		if err := k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap); err != nil {
			k.Logger(ctx).Error("could not close the monitored chain channel: %s", channelID)
		}
	}
}

func (k Keeper) IsMonitoredChanOpen(ctx sdk.Context, chainId string) bool {
	channelID, ok := k.GetChainToChannelMap(ctx, chainId)
	if !ok {
		return false
	}

	channel, ok := k.channelKeeper.GetChannel(ctx, k.GetPort(ctx), channelID)

	return ok && channel.State == channeltypes.OPEN
}
