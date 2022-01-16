package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	blockchain "github.com/cagnotteApp/Backend/blockchain"
	"github.com/cagnotteApp/Backend/domain"
	"github.com/cagnotteApp/Backend/repository"
	"github.com/cagnotteApp/Backend/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/go-amino"
)

const AddTxCollection = "AddTx"
const CreateCgtCollection = "CreateTx"

func ConfirmTx(cdc *amino.Codec, tx domain.TransactionDetails) ([]byte, error) {
	baseURL := utils.GetRestUrl()
	success, err := strconv.ParseBool(tx.Status)
	if err != nil {
		return nil, err
	}
	unsigedTx, err := blockchain.CreateConfirmTx(cdc, baseURL, tx, success)
	if err != nil {
		return nil, err
	}

	signedTx, err := blockchain.SignTx(unsigedTx, baseURL)

	if err != nil {
		return nil, err
	}

	newTx := domain.SignedTxModel{}
	newTx.Tx = signedTx
	newTx.Mode = "block"

	byteTx, err := cdc.MarshalJSON(newTx)
	if err != nil {
		return nil, err
	}
	txReceipt, err := utils.Request("POST", baseURL+"/txs", byteTx, nil)
	if err != nil {
		return nil, err
	}
	return txReceipt, nil

}

func getUserAccount(signature domain.Signature) (string, string) {

	var cdc = codec.New()

	codec.RegisterCrypto(cdc)

	userAccount := types.StdSignature{}

	ByteSignature, _ := cdc.MarshalJSON(signature)

	cdc.UnmarshalJSON(ByteSignature, &userAccount)

	PublicAddr, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, userAccount.PubKey)
	if err != nil {
		return "", ""
	}
	AccAddr := sdk.AccAddress(userAccount.PubKey.Address()).String()
	return AccAddr, PublicAddr

}

func VerifTransaction(repo repository.Repository, tx domain.TransactionDetails, db repository.DB) (bool, string, error) {
	userAddr, _ := getUserAccount(tx.Tx.Value.Signatures[0])

	userAccount, err := repo.FindUser(userAddr)
	if err != nil {
		return false, "", err
	}

	cagnotteName := tx.Tx.Value.Msg[0].Value.Name
	amount := tx.Tx.Value.Msg[0].Value.Bid
	if len(cagnotteName) == 0 || len(amount) == 0 {
		return false, "", errors.New("Cannot read the amount of money to validate ")
	}
	validFinancialTx, paiementToken, err := createFinancialTx(repo, *userAccount, tx)
	if !validFinancialTx {

		return false, "", nil
	}
	i := 0
	variable := os.Getenv("REPEAT")
	number, _ := strconv.Atoi(variable)
	if number == 0 {
		number = 50
	}
	execution := false
	for i < number {
		time.Sleep(5 * time.Second)
		execution, err = VerifPaiementExecution(repo, paiementToken)

		if err != nil {

			return false, "", nil
		}
		if execution {
			break
		}
		i++

	}

	return execution, paiementToken, nil
}

func initTransaction(cdc *amino.Codec, addrTo string) ([]byte, error) {

	baseURL := utils.GetRestUrl()

	unsigedTx, err := blockchain.CreateInitialTx(cdc, addrTo, baseURL)
	if err != nil {
		return nil, err
	}

	signedTx, err := blockchain.SignTx(unsigedTx, baseURL)

	if err != nil {
		return nil, err
	}

	newTx := domain.SignedTxModel{}
	newTx.Tx = signedTx
	newTx.Mode = "block"

	byteTx, err := cdc.MarshalJSON(newTx)
	if err != nil {
		return nil, err
	}
	txReceipt, err := utils.Request("POST", baseURL+"/txs", byteTx, nil)
	if err != nil {
		return nil, err
	}
	return txReceipt, nil

}

func GetValidConfirmations(repo repository.Repository, baseURL string) error {
	height, err := GetLatestBlock(baseURL)
	if err != nil {
		return err
	}
	latestheight, err := strconv.Atoi(height)
	if err != nil {
		return err
	}
	config, err := repo.GetConfig()
	minHeight := config.HeightValid + 1
	ConfirmTxs, err := GetTransactions(baseURL, "confirm_tx", minHeight, latestheight)

	if err != nil {
		return err
	}
	_, err = RegisterConfirmedTx(repo, ConfirmTxs, AddTxCollection)

	if err != nil {
		return err
	}
	config.HeightValid = latestheight
	_, err = repo.SaveConfig(config)
	if err != nil {
		return err
	}
	return nil

}

func GetTxFromBlockchain(repo repository.Repository, baseURL string, db repository.DB) error {
	height, err := GetLatestBlock(baseURL)
	if err != nil {
		return err
	}
	latestheight, err := strconv.Atoi(height)
	if err != nil {
		return err
	}
	config, err := repo.GetConfig()
	minHeight := config.HeightTx + 1

	CrtCgntTx, err := GetTransactions(baseURL, "CreateCagnotte", minHeight, latestheight)
	if err != nil {
		return err
	}
	AddCgntTx, err := GetTransactions(baseURL, "add_cagnotte", minHeight, latestheight)
	if err != nil {

		return err
	}

	_, err = RegisterTx(repo, CrtCgntTx, CreateCgtCollection)
	if err != nil {
		return err
	}

	_, err = RegisterTx(repo, AddCgntTx, AddTxCollection)
	if err != nil {
		return err
	}
	config.HeightTx = latestheight
	_, err = repo.SaveConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func GetTransactions(baseUrl string, operation string, minHeight int, maxheight int) ([]domain.TransactionDetails, error) {

	url := fmt.Sprintf("%s/txs?message.action=%s&tx.minheight=%d&tx.maxheight=%d", baseUrl, operation, minHeight, maxheight)
	data, err := utils.Request("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	transactions := domain.RegisteredTransaction{}
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return nil, err
	}
	transactionSet := []domain.TransactionDetails{}
	for _, tx := range transactions.Txs {
		transactionSet = append(transactionSet, domain.TransactionDetails{
			Height:     tx.Height,
			Hash:       tx.Txhash,
			Tx:         tx.Tx,
			Status:     "pending",
			Considered: false,
			Executed:   false,
		})
	}

	return transactionSet, nil
}

func GetLatestBlock(baseUrl string) (string, error) {
	url := fmt.Sprintf("%s/blocks/latest", baseUrl)
	data, err := utils.Request("GET", url, nil, nil)
	if err != nil {
		return "", err
	}
	block := domain.Block{}
	err = json.Unmarshal(data, &block)
	if err != nil {
		return "", err
	}
	return block.Block.Header.Height, nil
}
