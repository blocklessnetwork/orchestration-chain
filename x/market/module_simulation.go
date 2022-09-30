package market

import (
	"math/rand"

	"github.com/blocklessnetwork/orchestration-chain/testutil/sample"
	marketsimulation "github.com/blocklessnetwork/orchestration-chain/x/market/simulation"
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
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
	_ = marketsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSubmitOrder = "op_weight_msg_submit_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitOrder int = 100

	opWeightMsgSubmitCompletedOrder = "op_weight_msg_submit_completed_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitCompletedOrder int = 100

	opWeightMsgClaimOrderWork = "op_weight_msg_claim_order_work"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimOrderWork int = 100

	opWeightMsgRegisterHeadNode = "op_weight_msg_register_head_node"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterHeadNode int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	marketGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&marketGenesis)
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

	var weightMsgSubmitOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitOrder, &weightMsgSubmitOrder, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitOrder = defaultWeightMsgSubmitOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitOrder,
		marketsimulation.SimulateMsgSubmitOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitCompletedOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitCompletedOrder, &weightMsgSubmitCompletedOrder, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitCompletedOrder = defaultWeightMsgSubmitCompletedOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitCompletedOrder,
		marketsimulation.SimulateMsgSubmitCompletedOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimOrderWork int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimOrderWork, &weightMsgClaimOrderWork, nil,
		func(_ *rand.Rand) {
			weightMsgClaimOrderWork = defaultWeightMsgClaimOrderWork
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimOrderWork,
		marketsimulation.SimulateMsgClaimOrderWork(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterHeadNode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterHeadNode, &weightMsgRegisterHeadNode, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterHeadNode = defaultWeightMsgRegisterHeadNode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterHeadNode,
		marketsimulation.SimulateMsgRegisterHeadNode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
