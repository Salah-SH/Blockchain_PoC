package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgConfirmPendingTx struct {
	CagnotteName string         `json:"name"`
	Bid          int            `json:"bid"`
	Sender       sdk.AccAddress `json:"sender"`
	Success      bool           `json:"success"`
	User         sdk.AccAddress `json:"user"`
}

func NewMsgConfirmPendingTx(name string, bid int, sender sdk.AccAddress, succes bool, user sdk.AccAddress) MsgConfirmPendingTx {
	return MsgConfirmPendingTx{
		CagnotteName: name,
		Bid:          bid,
		Sender:       sender,
		Success:      succes,
		User:         user,
	}
}

func (msg MsgConfirmPendingTx) Route() string { return RouterKey }

func (msg MsgConfirmPendingTx) Type() string {
	return "confirm_tx"
}

func (msg MsgConfirmPendingTx) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender.String())
	}

	if len(msg.CagnotteName) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	if msg.Bid == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Amount to confirm cannot be 0")
	}
	if msg.User.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.User.String())
	}

	return nil
}

func (msg MsgConfirmPendingTx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgConfirmPendingTx) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
