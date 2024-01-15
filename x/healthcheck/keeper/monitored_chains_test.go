package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "healthcheck/testutil/keeper"
	"healthcheck/testutil/nullify"
	"healthcheck/x/healthcheck/keeper"
	"healthcheck/x/healthcheck/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMonitoredChains(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MonitoredChains {
	items := make([]types.MonitoredChains, n)
	for i := range items {
		items[i].ChainId = strconv.Itoa(i)

		keeper.SetMonitoredChains(ctx, items[i])
	}
	return items
}

func TestMonitoredChainsGet(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	items := createNMonitoredChains(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMonitoredChains(ctx,
			item.ChainId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMonitoredChainsRemove(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	items := createNMonitoredChains(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMonitoredChains(ctx,
			item.ChainId,
		)
		_, found := keeper.GetMonitoredChains(ctx,
			item.ChainId,
		)
		require.False(t, found)
	}
}

func TestMonitoredChainsGetAll(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	items := createNMonitoredChains(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMonitoredChains(ctx)),
	)
}
