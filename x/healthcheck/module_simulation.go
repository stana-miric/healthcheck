package healthcheck

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"healthcheck/testutil/sample"
	healthchecksimulation "healthcheck/x/healthcheck/simulation"
	"healthcheck/x/healthcheck/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = healthchecksimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateMonitoredChains = "op_weight_msg_monitored_chains"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMonitoredChains int = 100

	opWeightMsgUpdateMonitoredChains = "op_weight_msg_monitored_chains"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMonitoredChains int = 100

	opWeightMsgDeleteMonitoredChains = "op_weight_msg_monitored_chains"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMonitoredChains int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	healthcheckGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		MonitoredChainsList: []types.MonitoredChains{
			{
				Creator: sample.AccAddress(),
				ChainId: "0",
			},
			{
				Creator: sample.AccAddress(),
				ChainId: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&healthcheckGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMonitoredChains int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMonitoredChains, &weightMsgCreateMonitoredChains, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMonitoredChains = defaultWeightMsgCreateMonitoredChains
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMonitoredChains,
		healthchecksimulation.SimulateMsgCreateMonitoredChains(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMonitoredChains int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMonitoredChains, &weightMsgUpdateMonitoredChains, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMonitoredChains = defaultWeightMsgUpdateMonitoredChains
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMonitoredChains,
		healthchecksimulation.SimulateMsgUpdateMonitoredChains(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMonitoredChains int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMonitoredChains, &weightMsgDeleteMonitoredChains, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMonitoredChains = defaultWeightMsgDeleteMonitoredChains
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMonitoredChains,
		healthchecksimulation.SimulateMsgDeleteMonitoredChains(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
