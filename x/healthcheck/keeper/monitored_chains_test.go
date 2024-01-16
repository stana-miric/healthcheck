package keeper_test

import (
	"strconv"
	"testing"

	keepertest "healthcheck/testutil/keeper"
	"healthcheck/testutil/nullify"
	"healthcheck/x/healthcheck/keeper"
	"healthcheck/x/healthcheck/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMonitoredChain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MonitoredChain {
	items := make([]types.MonitoredChain, n)
	for i := range items {
		items[i].ChainId = strconv.Itoa(i)

		keeper.SetMonitoredChain(ctx, items[i])
	}
	return items
}

func TestMonitoredChainGet(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	items := createNMonitoredChain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMonitoredChain(ctx,
			item.ChainId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMonitoredChainRemove(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	items := createNMonitoredChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMonitoredChain(ctx,
			item.ChainId,
		)
		_, found := keeper.GetMonitoredChain(ctx,
			item.ChainId,
		)
		require.False(t, found)
	}
}

func TestMonitoredChainGetAll(t *testing.T) {
	keeper, ctx := keepertest.HealthcheckKeeper(t)
	items := createNMonitoredChain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMonitoredChain(ctx)),
	)
}
