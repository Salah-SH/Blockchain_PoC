package cagnotte

import (
	"fmt"

	"github.com/cagnotteApp/x/cagnotte/keeper"
	"github.com/cagnotteApp/x/cagnotte/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgAddCagnotte:
			return handleMsgAddCagnotte(ctx, k, msg)
		case types.MsgCreateCagnotte:
			return handleMsgCreateCagnotte(ctx, k, msg)
		case types.MsgCloseCagnotte:
			return handleMsgCloseCagnotte(ctx, k, msg)
		case types.MsgConfirmPendingTx:
			return handleMsgConfirmPendingTx(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
