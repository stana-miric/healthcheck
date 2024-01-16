package keeper

import (
	"healthcheck/x/healthcheck/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetMonitoredChain set a specific MonitoredChain in the store from its index
func (k Keeper) SetMonitoredChain(ctx sdk.Context, MonitoredChain types.MonitoredChain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainKeyPrefix))
	b := k.cdc.MustMarshal(&MonitoredChain)
	store.Set(types.MonitoredChainKey(
		MonitoredChain.ChainId,
	), b)
}

// GetMonitoredChain returns a MonitoredChain from its index
func (k Keeper) GetMonitoredChain(
	ctx sdk.Context,
	chainId string,

) (val types.MonitoredChain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainKeyPrefix))

	b := store.Get(types.MonitoredChainKey(
		chainId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMonitoredChain removes a MonitoredChain from the store
func (k Keeper) RemoveMonitoredChain(
	ctx sdk.Context,
	chainId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainKeyPrefix))
	store.Delete(types.MonitoredChainKey(
		chainId,
	))
}

// GetAllMonitoredChain returns all MonitoredChain
func (k Keeper) GetAllMonitoredChain(ctx sdk.Context) (list []types.MonitoredChain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MonitoredChain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetChainToChannel sets the mapping from a monitored chainID to the channel ID for that monitored chain.
func (k Keeper) SetChainToChannel(ctx sdk.Context, chainID, channelID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainToChannelKeyPrefix))
	store.Set(types.ChainToChannelKey(
		chainID,
	), []byte(channelID))

	store.Set(types.ChainToChannelKey(chainID), []byte(channelID))
}

// GetChainToChannel gets the channelID for the given moniotred chainID
func (k Keeper) GetChainToChannel(ctx sdk.Context, chainID string) (string, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainToChannelKeyPrefix))

	bz := store.Get(types.ChainToChannelKey(
		chainID,
	))
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

func (k Keeper) InitializeMonitoredChain(ctx sdk.Context, connectionID string, timeoutInterval, updateInterval uint32) error {

	chainID, err := k.GetClientChainIdFromConnection(ctx, connectionID)
	if err != nil {
		return err
	}

	chain, found := k.GetMonitoredChain(ctx, chainID)

	if !found {
		return sdkerrors.Wrapf(types.ErrUnregisteredChain, "chain-id: %s", chainID)
	}

	if chain.ConnectionId != connectionID {
		return sdkerrors.Wrapf(types.ErrInvalidConnection, "connection-id: %s", connectionID)
	}

	var updatedChain = types.MonitoredChain{
		Creator:         chain.Creator,
		ChainId:         chain.ChainId,
		ConnectionId:    chain.ConnectionId,
		TimeoutInterval: timeoutInterval,
		UpdateInterval:  updateInterval,
		Status: &types.MonitoredChainStatus{
			Status: "inactive",
		},
	}

	k.SetMonitoredChain(ctx, updatedChain)

	return nil
}
