package types_test

import (
	"testing"

	"healthcheck/x/healthcheck/types"
	commonTypes "healthcheck/x/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: commonTypes.HealthcheckPortID,
				MonitoredChainList: []types.MonitoredChain{
					{
						ChainId: "0",
					},
					{
						ChainId: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated MonitoredChain",
			genState: &types.GenesisState{
				MonitoredChainList: []types.MonitoredChain{
					{
						ChainId: "0",
					},
					{
						ChainId: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
