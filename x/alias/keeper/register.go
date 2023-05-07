package keeper

import (
	"units/x/alias/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"
)

func (k Keeper) AliasIsFree(ctx sdk.Context, name string) bool {
	allAliases := k.GetAllAliases(ctx)
	for _, a := range allAliases {
		if slices.Contains(a.Names, name) {
			return false
		}
	}
	return true
}

func (k Keeper) Register(ctx sdk.Context, owner string, name string) error {

	// 0. check if owner is a valid address
	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return err
	}

	// 1. check alias is free
	if !k.AliasIsFree(ctx, name) {
		return types.ErrAliasNotFree
	}

	// 2. check max entries
	aliases, found := k.GetAliases(ctx, owner)
	if len(aliases.Names) >= int(k.GetParams(ctx).MaxEntries) {
		return types.ErrMaxEntries
	}

	// initialize if not found
	if !found {
		aliases.Address = owner
	}

	// append nickname
	aliases.Names = append(aliases.Names, name)

	// set aliases
	k.SetAliases(ctx, aliases)

	return nil
}
