package cagnotte

import (
	"github.com/cagnotteApp/x/cagnotte/keeper"
	"github.com/cagnotteApp/x/cagnotte/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Handle a message to participate in a cagnotte
func handleMsgAddCagnotte(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAddCagnotte) (*sdk.Result, error) {
	exist := keeper.Exists(ctx, msg.Name)
	if !exist {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidSequence, "The cagnotte  is not found in the db")
	}
	active := keeper.IsActive(ctx, msg.Name)
	if !active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "This cagnotte is  closed")
	}

	keeper.AddParticipator(ctx, msg.Name, msg.Participator)
	keeper.AddPendingAmount(ctx, msg.Name, msg.Participator, msg.Bid)
	return &sdk.Result{}, nil
}
