package keeper_test

import (
	"context"
	"testing"

	keepertest "cifer/testutil/keeper"
	"cifer/testutil/nullify"
	"cifer/x/cifer/keeper"
	"cifer/x/cifer/types"

	"github.com/stretchr/testify/require"
)

func createNMintdata(keeper keeper.Keeper, ctx context.Context, n int) []types.Mintdata {
	items := make([]types.Mintdata, n)
	for i := range items {
		items[i].Id = keeper.AppendMintdata(ctx, items[i])
	}
	return items
}

func TestMintdataGet(t *testing.T) {
	keeper, ctx := keepertest.CiferKeeper(t)
	items := createNMintdata(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMintdata(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMintdataRemove(t *testing.T) {
	keeper, ctx := keepertest.CiferKeeper(t)
	items := createNMintdata(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMintdata(ctx, item.Id)
		_, found := keeper.GetMintdata(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMintdataGetAll(t *testing.T) {
	keeper, ctx := keepertest.CiferKeeper(t)
	items := createNMintdata(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMintdata(ctx)),
	)
}

func TestMintdataCount(t *testing.T) {
	keeper, ctx := keepertest.CiferKeeper(t)
	items := createNMintdata(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMintdataCount(ctx))
}
