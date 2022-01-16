package cagnotte

import (
	"github.com/cagnotteApp/x/cagnotte/keeper"
	"github.com/cagnotteApp/x/cagnotte/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgCloseCagnotte(ctx sdk.Context, k keeper.Keeper, msg types.MsgCloseCagnotte) (*sdk.Result, error) {
	cagnotte, err := k.GetCagnotte(ctx, msg.Name)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidSequence, "The cagnotte  is not found in the db")

	}

	if !msg.Executer.Equals(cagnotte.Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	active := k.IsActive(ctx, msg.Name)
	if !active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "This cagnotte is already closed")
	}
	err = k.CloseCagnotte(ctx, msg.Name)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The cagnotte cannot be closed, please do execute pending transaction before closure!")
	}
	return &sdk.Result{}, nil
}
