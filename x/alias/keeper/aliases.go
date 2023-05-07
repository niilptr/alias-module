package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"units/x/alias/types"
)

// SetAliases set a specific aliases in the store from its index
func (k Keeper) SetAliases(ctx sdk.Context, aliases types.Aliases) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AliasesKeyPrefix))
	b := k.cdc.MustMarshal(&aliases)
	store.Set(types.AliasesKey(
		aliases.Address,
	), b)
}

// GetAliases returns a aliases from its index
func (k Keeper) GetAliases(
	ctx sdk.Context,
	address string,

) (val types.Aliases, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AliasesKeyPrefix))

	b := store.Get(types.AliasesKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAliases removes a aliases from the store
func (k Keeper) RemoveAliases(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AliasesKeyPrefix))
	store.Delete(types.AliasesKey(
		address,
	))
}

// GetAllAliases returns all aliases
func (k Keeper) GetAllAliases(ctx sdk.Context) (list []types.Aliases) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AliasesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Aliases
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
