package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCloseCagnotte defines the BuyName message
type MsgCloseCagnotte struct {
	Name     string         `json:"name"`
	Executer sdk.AccAddress `json:"executor"`
}

// NewMsgCloseCagnotte is the constructor function for MsgCloseCagnotte
func NewMsgCloseCagnotte(name string, account sdk.AccAddress) MsgCloseCagnotte {
	return MsgCloseCagnotte{
		Name:     name,
		Executer: account,
	}
}

// Route should return the name of the module
func (msg MsgCloseCagnotte) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCloseCagnotte) Type() string { return "close-cagnotte" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCloseCagnotte) ValidateBasic() error {
	if msg.Executer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Executer.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	// if !msg.Bid.IsAllPositive() {
	// 	return sdkerrors.ErrInsufficientFunds
	// }
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCloseCagnotte) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCloseCagnotte) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Executer}
}
