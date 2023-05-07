package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"units/x/alias/types"
)

func (k msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterResponse{}, nil
}
