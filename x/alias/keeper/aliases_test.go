package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "units/testutil/keeper"
	"units/testutil/nullify"
	"units/x/alias/keeper"
	"units/x/alias/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAliases(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Aliases {
	items := make([]types.Aliases, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetAliases(ctx, items[i])
	}
	return items
}

func TestAliasesGet(t *testing.T) {
	keeper, ctx := keepertest.AliasKeeper(t)
	items := createNAliases(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAliases(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAliasesRemove(t *testing.T) {
	keeper, ctx := keepertest.AliasKeeper(t)
	items := createNAliases(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAliases(ctx,
			item.Address,
		)
		_, found := keeper.GetAliases(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestAliasesGetAll(t *testing.T) {
	keeper, ctx := keepertest.AliasKeeper(t)
	items := createNAliases(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAliases(ctx)),
	)
}
