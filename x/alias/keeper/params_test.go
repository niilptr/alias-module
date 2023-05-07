package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "units/testutil/keeper"
	"units/x/alias/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AliasKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
