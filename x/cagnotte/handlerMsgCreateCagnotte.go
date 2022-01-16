package cagnotte

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cagnotteApp/x/cagnotte/keeper"
	"github.com/cagnotteApp/x/cagnotte/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgCreateCagnotte(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateCagnotte) (*sdk.Result, error) {
	err := k.CreateCagnotte(ctx, msg.Name, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, " This name is already used")
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
