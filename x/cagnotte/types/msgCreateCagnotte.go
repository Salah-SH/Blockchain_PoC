package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgCreateCagnotte struct {
	Name    string         `json:"name"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgCreateCagnotte(name string, creator sdk.AccAddress) MsgCreateCagnotte {

	return MsgCreateCagnotte{
		Name:    name,
		Creator: creator,
	}
}

func (msg MsgCreateCagnotte) Route() string {
	return RouterKey
}

func (msg MsgCreateCagnotte) Type() string {
	return "CreateCagnotte"
}

func (msg MsgCreateCagnotte) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	return nil
}

func (msg MsgCreateCagnotte) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateCagnotte) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
