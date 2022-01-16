package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgAddCagnotte struct {
	Name         string         `json:"name"`
	Bid          int            `json:"bid"`
	Participator sdk.AccAddress `json:"participator"`
}

func NewMsgAddCagnotte(name string, bid int, participator sdk.AccAddress) MsgAddCagnotte {
	return MsgAddCagnotte{
		Name:         name,
		Bid:          bid,
		Participator: participator,
	}
}

// Route should return the name of the module
func (msg MsgAddCagnotte) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddCagnotte) Type() string {

	return "add_cagnotte"
}

// ValidateBasic runs stateless checks on the message
func (msg MsgAddCagnotte) ValidateBasic() error {
	if msg.Participator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Participator.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddCagnotte) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAddCagnotte) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Participator}
}
