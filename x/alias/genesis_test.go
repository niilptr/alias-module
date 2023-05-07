package alias_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "units/testutil/keeper"
	"units/testutil/nullify"
	"units/x/alias"
	"units/x/alias/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AliasesList: []types.Aliases{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AliasKeeper(t)
	alias.InitGenesis(ctx, *k, genesisState)
	got := alias.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AliasesList, got.AliasesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
