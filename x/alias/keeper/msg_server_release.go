package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"units/x/alias/types"
)

func (k msgServer) Release(goCtx context.Context, msg *types.MsgRelease) (*types.MsgReleaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgReleaseResponse{}, nil
}
