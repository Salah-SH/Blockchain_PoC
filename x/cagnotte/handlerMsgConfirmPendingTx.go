package cagnotte

import (
	"errors"

	"github.com/cagnotteApp/x/cagnotte/keeper"
	"github.com/cagnotteApp/x/cagnotte/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgConfirmPendingTx(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgConfirmPendingTx) (*sdk.Result, error) {

	exist := keeper.Exists(ctx, msg.CagnotteName)
	if !exist {
		return nil, errors.New("the cagnotte is not found")
	}

	active := keeper.IsActive(ctx, msg.CagnotteName)
	if !active {

		return nil, errors.New("The cagnotte %s is closed")
	}
	allowed, err := keeper.Allowed(ctx, msg.Sender)
	if err != nil {

		return nil, err
	}
	if !allowed {

		return nil, errors.New("th is user is not allowed to confirm Tx")
	}
	err = keeper.AddAmount(ctx, msg.CagnotteName, msg.Bid, msg.Sender, msg.Success, msg.User)
	if err != nil {

		return nil, err
	}
	return &sdk.Result{}, nil
}
