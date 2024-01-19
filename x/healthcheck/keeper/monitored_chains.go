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

// SetChainToChannelMap sets the mapping from a monitored chainID to the channel ID for that monitored chain.
func (k Keeper) SetChainToChannelMap(ctx sdk.Context, chainID, channelID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainToChannelKeyPrefix))
	store.Set(types.ChainToChannelKey(
		chainID,
	), []byte(channelID))

	store.Set(types.ChainToChannelKey(chainID), []byte(channelID))
}

// GetChainToChannelMap gets the channelID for the given moniotred chainID
func (k Keeper) GetChainToChannelMap(ctx sdk.Context, chainID string) (string, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainToChannelKeyPrefix))

	bz := store.Get(types.ChainToChannelKey(
		chainID,
	))
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

// RemoveChanFromChainToChannelMap removes the channelID for the given moniotred chainID
func (k Keeper) RemoveChanFromChainToChannelMap(ctx sdk.Context, chainID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainToChannelKeyPrefix))

	store.Delete(types.ChainToChannelKey(chainID))
}

func (k Keeper) InitializeMonitoredChain(ctx sdk.Context, connectionID string, timeoutInterval, updateInterval uint64) error {

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

	chain.TimeoutInterval = timeoutInterval
	chain.UpdateInterval = updateInterval
	chain.Status = &types.MonitoredChainStatus{
		Status: string(types.Inactive),
	}

	k.SetMonitoredChain(ctx, chain)

	return nil
}
