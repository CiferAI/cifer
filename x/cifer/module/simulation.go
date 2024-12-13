package cifer

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"cifer/testutil/sample"
	cifersimulation "cifer/x/cifer/simulation"
	"cifer/x/cifer/types"
)

// avoid unused import issue
var (
	_ = cifersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateMintdata = "op_weight_msg_mintdata"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMintdata int = 100

	opWeightMsgUpdateMintdata = "op_weight_msg_mintdata"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMintdata int = 100

	opWeightMsgDeleteMintdata = "op_weight_msg_mintdata"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMintdata int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ciferGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MintdataList: []types.Mintdata{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MintdataCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ciferGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMintdata int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMintdata, &weightMsgCreateMintdata, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMintdata = defaultWeightMsgCreateMintdata
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMintdata,
		cifersimulation.SimulateMsgCreateMintdata(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMintdata int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMintdata, &weightMsgUpdateMintdata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMintdata = defaultWeightMsgUpdateMintdata
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMintdata,
		cifersimulation.SimulateMsgUpdateMintdata(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMintdata int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMintdata, &weightMsgDeleteMintdata, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMintdata = defaultWeightMsgDeleteMintdata
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMintdata,
		cifersimulation.SimulateMsgDeleteMintdata(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMintdata,
			defaultWeightMsgCreateMintdata,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cifersimulation.SimulateMsgCreateMintdata(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMintdata,
			defaultWeightMsgUpdateMintdata,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cifersimulation.SimulateMsgUpdateMintdata(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteMintdata,
			defaultWeightMsgDeleteMintdata,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cifersimulation.SimulateMsgDeleteMintdata(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
