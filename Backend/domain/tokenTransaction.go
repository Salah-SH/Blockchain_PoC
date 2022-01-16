package domain

import "github.com/cosmos/cosmos-sdk/x/auth/types"

type InitTx struct {
	Basereq BaseReq `json:"base_req"`
	Amount  []Coin  `json:"amount"`
}
type confirmTxModel struct {
	Basereq          BaseReq `json:"base_req"`
	CagnotteName     string  `json:"cagnotte_name"`
	Amount           string  `json:"amount"`
	Result           bool    `json:"result"`
	ParticipatorAddr string  `json:"participator"`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
type BaseReq struct {
	From          string `json:"from"`
	Memo          string `json:"memo"`
	ChainID       string `json:"chain_id"`
	AccountNumber string `json:"account_number"`
	Sequence      string `json:"sequence"`
	Gas           string `json:"gas"`
	GasAdjustment string `json:"gas_adjustment"`
	Simulate      bool   `json:"simulate"`
}

type SignedTxModel struct {
	Type string      `json:"type"`
	Tx   types.StdTx `json:"tx"`
	Mode string      `json:"mode,omitempty" bson:"omitempty"`
}

func InitialiseSendTx(addr string, chainId string, accountNumber string, sequence string, amount string) InitTx {
	return InitTx{
		Basereq: BaseReq{
			From:          addr,
			ChainID:       chainId,
			Sequence:      sequence,
			AccountNumber: accountNumber,
			Gas:           "200000",
			Simulate:      false,
		},
		Amount: []Coin{
			{
				Denom:  "nametoken",
				Amount: amount,
			},
		},
	}
}

func InitialiseConfirmTx(FromAddr string, chainId string, accountNumber string, sequence string, amount string, cagnotteName string, participatorAddr string, result bool) confirmTxModel {
	return confirmTxModel{
		Basereq: BaseReq{
			From:          FromAddr,
			ChainID:       chainId,
			Sequence:      sequence,
			AccountNumber: accountNumber,
			Gas:           "200000",
			Simulate:      false,
		},
		Amount:           amount,
		CagnotteName:     cagnotteName,
		Result:           result,
		ParticipatorAddr: participatorAddr,
	}
}
