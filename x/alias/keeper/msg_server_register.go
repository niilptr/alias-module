package keeper

import (
	"context"

	"units/x/alias/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (ms msgServer) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := ms.Keeper.Register(ctx, msg.Creator, msg.Name)
	if err != nil {
		return &types.MsgRegisterResponse{}, nil
	}

	return &types.MsgRegisterResponse{}, nil
}
