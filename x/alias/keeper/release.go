package keeper

import (
	"units/x/alias/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"
)

func (k Keeper) Release(ctx sdk.Context, owner string, name string) error {

	// 0. check if owner is a valid address
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return err
	}

	// 1. check if alias is found
	aliases, found := k.GetAliases(ctx, owner)
	if !found || !slices.Contains(aliases.Names, name) {
		return types.ErrAliasNotFound
	}

	// remove alias from saved
	idx := slices.Index(aliases.Names, name)
	aliases.Names = slices.Delete(aliases.Names, idx, idx+1)

	// set aliases
	k.SetAliases(ctx, aliases)

	return nil
}
