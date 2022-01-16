package domain

import (
	"time"
)

type AccountStruct struct {
	Height string  `json:"height"`
	Result account `json:"result"`
}

type account struct {
	Type  string `json:"type"`
	Value struct {
		Address string `json:"address"`
		Coins   []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"coins"`
		PublicKey struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"public_key"`
		AccountNumber string `json:"account_number"`
		Sequence      string `json:"sequence"`
	} `json:"value"`
}

type RegisteredTransaction struct {
	TotalCount string `json:"total_count"`
	Count      string `json:"count"`
	PageNumber string `json:"page_number"`
	PageTotal  string `json:"page_total"`
	Limit      string `json:"limit"`
	Txs        []struct {
		Height    string    `json:"height"`
		Txhash    string    `json:"txhash"`
		RawLog    string    `json:"raw_log"`
		GasWanted string    `json:"gas_wanted"`
		GasUsed   string    `json:"gas_used"`
		Tx        TxModel   `json:"tx"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"txs"`
}
type TransactionDetails struct {
	Height              string  `json:"height"`
	Hash                string  `json:"hash" bson:"_id"`
	Tx                  TxModel `json:"tx"`
	Considered          bool    `json:"considered"`
	Status              string  `json:"status"`
	Executed            bool    `json:"executed"`           // the validation transaction executed by the admin
	ConfimationonLedger bool    `json:"confimationonedger"` // whether the confirmation transaction is validated on the blck
	ConfirmTxHash       string  `json:"Confirmtxhash"`      // the hash of the transaction executed by admin to validate the tx
}

type TxModel struct {
	Type  string `json:"type"`
	Value struct {
		Msg []struct {
			Type  string `json:"type"`
			Value struct {
				Name         string `json:"name"`
				Creator      string `json:"creator"  `
				Bid          string `json:"bid"`
				Participator string `json:"participator"`
			} `json:"value"`
		} `json:"msg"`
		Fee struct {
			Amount []struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"amount"`
			Gas string `json:"gas"`
		} `json:"fee"`
		Signatures []Signature `json:"signatures"`
		Memo       string      `json:"memo"`
	} `json:"value"`
	Mode string `json:"mode,omitempty" bson:"omitempty"`
}
type Signature struct {
	PubKey    PublicKey `json:"pub_key"`
	Signature string    `json:"signature"`
}
type PublicKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Config struct {
	HeightTx    int `json:"heighttx"`
	HeightValid int `json:"heightvalid"`
}

type TxReceipt struct {
	Height string `json:"height"`
	Txhash string `json:"txhash"`
	Code   int    `json:"code"`
	RawLog string `json:"raw_log"`
}
