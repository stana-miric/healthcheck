package keeper

import (
	"healthcheck/x/healthcheck/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) StartTrackingMonitoredChain(connectionID string, timeoutInterval, updateInterval uint32) error {

	// if err := k.UpdateRegisteredChainInterval(connectionID, timeoutInterval, updateInterval); err != nil {
	// 	return err
	// }

	// if err := k.SetMonitoredChainEntry(); err != nil {
	// 	return err
	// }

	return nil
}
