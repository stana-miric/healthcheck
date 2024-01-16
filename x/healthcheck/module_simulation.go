package healthcheck

import (
	"math/rand"

	"healthcheck/testutil/sample"
	healthchecksimulation "healthcheck/x/healthcheck/simulation"
	"healthcheck/x/healthcheck/types"
	commonTypes "healthcheck/x/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
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
	opWeightMsgCreateMonitoredChain = "op_weight_msg_monitored_chains"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMonitoredChain int = 100

	opWeightMsgUpdateMonitoredChain = "op_weight_msg_monitored_chains"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMonitoredChain int = 100

	opWeightMsgDeleteMonitoredChain = "op_weight_msg_monitored_chains"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMonitoredChain int = 100

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
		PortId: commonTypes.HealthcheckPortID,
		MonitoredChainList: []types.MonitoredChain{
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

	var weightMsgCreateMonitoredChain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMonitoredChain, &weightMsgCreateMonitoredChain, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMonitoredChain = defaultWeightMsgCreateMonitoredChain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMonitoredChain,
		healthchecksimulation.SimulateMsgCreateMonitoredChain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMonitoredChain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMonitoredChain, &weightMsgUpdateMonitoredChain, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMonitoredChain = defaultWeightMsgUpdateMonitoredChain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMonitoredChain,
		healthchecksimulation.SimulateMsgUpdateMonitoredChain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMonitoredChain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMonitoredChain, &weightMsgDeleteMonitoredChain, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMonitoredChain = defaultWeightMsgDeleteMonitoredChain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMonitoredChain,
		healthchecksimulation.SimulateMsgDeleteMonitoredChain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
