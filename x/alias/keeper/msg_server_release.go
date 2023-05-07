package keeper

import (
	"context"

	"units/x/alias/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (ms msgServer) Release(goCtx context.Context, msg *types.MsgRelease) (*types.MsgReleaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := ms.Keeper.Release(ctx, msg.Creator, msg.Name)
	if err != nil {
		return &types.MsgReleaseResponse{}, err
	}

	return &types.MsgReleaseResponse{}, nil
}
