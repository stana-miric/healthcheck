package keeper

import (
	"testing"

	"healthcheck/x/healthcheck/keeper"
	"healthcheck/x/healthcheck/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// healthcheckChannelKeeper is a stub of cosmosibckeeper.ChannelKeeper.
type healthcheckChannelKeeper struct{}

func (healthcheckChannelKeeper) GetChannel(ctx sdk.Context, portID, channelID string) (channeltypes.Channel, bool) {
	return channeltypes.Channel{}, false
}

func (healthcheckChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return 0, false
}

func (healthcheckChannelKeeper) SendPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	return 0, nil
}

func (healthcheckChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	return nil
}

func (healthcheckChannelKeeper) GetChannelClientState(ctx sdk.Context, portID, channelID string) (string, exported.ClientState, error) {
	return "", nil, nil
}

// healthcheckportKeeper is a stub of cosmosibckeeper.PortKeeper
type healthcheckPortKeeper struct{}

func (healthcheckPortKeeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return &capabilitytypes.Capability{}
}

// healthcheckConnectionKeeper is a stub of cosmosibckeeper.ConnectionKeeper
type healthcheckConnectionKeeper struct{}

func (healthcheckConnectionKeeper) GetConnection(ctx sdk.Context, connectionID string) (conntypes.ConnectionEnd, bool) {
	return conntypes.ConnectionEnd{}, false
}

// healthcheckClientKeeper is a stub of cosmosibckeeper.Client
type healthcheckClientKeeper struct{}

func (healthcheckClientKeeper) GetClientState(ctx sdk.Context, clientID string) (exported.ClientState, bool) {
	return nil, false
}

func HealthcheckKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"HealthcheckParams",
	)
	k := keeper.NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		healthcheckChannelKeeper{},
		healthcheckPortKeeper{},
		healthcheckConnectionKeeper{},
		healthcheckClientKeeper{},
		capabilityKeeper.ScopeToModule("HealthcheckScopedKeeper"),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
