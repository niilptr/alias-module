package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"units/x/alias/types"
)

func (k Keeper) AliasesAll(goCtx context.Context, req *types.QueryAllAliasesRequest) (*types.QueryAllAliasesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var aliasess []types.Aliases
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	aliasesStore := prefix.NewStore(store, types.KeyPrefix(types.AliasesKeyPrefix))

	pageRes, err := query.Paginate(aliasesStore, req.Pagination, func(key []byte, value []byte) error {
		var aliases types.Aliases
		if err := k.cdc.Unmarshal(value, &aliases); err != nil {
			return err
		}

		aliasess = append(aliasess, aliases)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAliasesResponse{Aliases: aliasess, Pagination: pageRes}, nil
}

func (k Keeper) Aliases(goCtx context.Context, req *types.QueryGetAliasesRequest) (*types.QueryGetAliasesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetAliases(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAliasesResponse{Aliases: val}, nil
}
