package cifer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cifer/x/cifer/keeper"
	"cifer/x/cifer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the mintdata
	for _, elem := range genState.MintdataList {
		k.SetMintdata(ctx, elem)
	}

	// Set mintdata count
	k.SetMintdataCount(ctx, genState.MintdataCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MintdataList = k.GetAllMintdata(ctx)
	genesis.MintdataCount = k.GetMintdataCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
