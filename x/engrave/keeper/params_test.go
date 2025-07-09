package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "cifer/testutil/keeper"
	"cifer/x/engrave/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.EngraveKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
