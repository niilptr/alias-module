package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"units/x/alias/keeper"
	"units/x/alias/types"
)

func SimulateMsgRelease(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRelease{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Release simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Release simulation not implemented"), nil, nil
	}
}
