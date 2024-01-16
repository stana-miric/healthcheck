package healthcheck_test

import (
	"testing"

	keepertest "healthcheck/testutil/keeper"
	"healthcheck/testutil/nullify"
	"healthcheck/x/healthcheck"
	"healthcheck/x/healthcheck/types"
	commonTypes "healthcheck/x/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: commonTypes.HealthcheckPortID,
		MonitoredChainList: []types.MonitoredChain{
			{
				ChainId: "0",
			},
			{
				ChainId: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HealthcheckKeeper(t)
	healthcheck.InitGenesis(ctx, *k, genesisState)
	got := healthcheck.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.MonitoredChainList, got.MonitoredChainList)
	// this line is used by starport scaffolding # genesis/test/assert
}
