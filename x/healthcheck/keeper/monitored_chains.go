package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"healthcheck/x/healthcheck/types"
)

// SetMonitoredChains set a specific monitoredChains in the store from its index
func (k Keeper) SetMonitoredChains(ctx sdk.Context, monitoredChains types.MonitoredChains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainsKeyPrefix))
	b := k.cdc.MustMarshal(&monitoredChains)
	store.Set(types.MonitoredChainsKey(
		monitoredChains.ChainId,
	), b)
}

// GetMonitoredChains returns a monitoredChains from its index
func (k Keeper) GetMonitoredChains(
	ctx sdk.Context,
	chainId string,

) (val types.MonitoredChains, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainsKeyPrefix))

	b := store.Get(types.MonitoredChainsKey(
		chainId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMonitoredChains removes a monitoredChains from the store
func (k Keeper) RemoveMonitoredChains(
	ctx sdk.Context,
	chainId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainsKeyPrefix))
	store.Delete(types.MonitoredChainsKey(
		chainId,
	))
}

// GetAllMonitoredChains returns all monitoredChains
func (k Keeper) GetAllMonitoredChains(ctx sdk.Context) (list []types.MonitoredChains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoredChainsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MonitoredChains
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
