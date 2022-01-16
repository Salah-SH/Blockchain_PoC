package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type hash []byte
type Cagnotte struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Name  string         `json:"name" yaml:"name"`
	//	Value string         `json:"value" yaml:"value"`
	Participators       []sdk.AccAddress
	Status              bool
	Amount              int            `json:"amount" yaml:"amount"`
	PendingTransactions []Participator `json:"pendingamount" yaml:"pendingamount"`
	InvalidTransactions []Participator `json:"invalidtxs" yaml:"invalidtxs"`
	ValidTransactions   []Participator `json:"validtxs" yaml:"validtxs"`
}

type Participator struct {
	User   sdk.AccAddress `json:"user" yaml:"user"`
	Amount int            `json:"amount" yaml:"amount"`
}
type AdminAddr struct {
	Address sdk.AccAddress `json:"admin_address"`
}
