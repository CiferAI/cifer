package cifer_test

import (
	"testing"

	keepertest "cifer/testutil/keeper"
	"cifer/testutil/nullify"
	"cifer/x/cifer/module"
	"cifer/x/cifer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CiferKeeper(t)
	cifer.InitGenesis(ctx, k, genesisState)
	got := cifer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
