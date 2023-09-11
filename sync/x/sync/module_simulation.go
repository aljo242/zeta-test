package sync

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"sync/testutil/sample"
	syncsimulation "sync/x/sync/simulation"
	"sync/x/sync/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = syncsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateHeader = "op_weight_msg_header"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHeader int = 100

	opWeightMsgUpdateHeader = "op_weight_msg_header"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHeader int = 100

	opWeightMsgDeleteHeader = "op_weight_msg_header"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHeader int = 100

	opWeightMsgAdmin = "op_weight_msg_admin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAdmin int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	syncGenesis := types.GenesisState{
		Params: types.DefaultParams(),

		HeaderList: []types.Header{
			{
				BlockID: 0,
			},
			{
				BlockID: 1,
			},
		},
		HeaderCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&syncGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateHeader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHeader, &weightMsgCreateHeader, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHeader = defaultWeightMsgCreateHeader
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHeader,
		syncsimulation.SimulateMsgCreateHeader(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHeader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateHeader, &weightMsgUpdateHeader, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHeader = defaultWeightMsgUpdateHeader
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHeader,
		syncsimulation.SimulateMsgUpdateHeader(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHeader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHeader, &weightMsgDeleteHeader, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHeader = defaultWeightMsgDeleteHeader
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHeader,
		syncsimulation.SimulateMsgDeleteHeader(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAdmin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAdmin, &weightMsgAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgAdmin = defaultWeightMsgAdmin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAdmin,
		syncsimulation.SimulateMsgAdmin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateHeader,
			defaultWeightMsgCreateHeader,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				syncsimulation.SimulateMsgCreateHeader(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateHeader,
			defaultWeightMsgUpdateHeader,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				syncsimulation.SimulateMsgUpdateHeader(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteHeader,
			defaultWeightMsgDeleteHeader,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				syncsimulation.SimulateMsgDeleteHeader(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAdmin,
			defaultWeightMsgAdmin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				syncsimulation.SimulateMsgAdmin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
