package listener

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/cagnotteApp/Backend/domain"
	"github.com/cagnotteApp/Backend/repository"
	"github.com/cagnotteApp/Backend/service"
	"github.com/cagnotteApp/Backend/utils"
	"github.com/jasonlvhit/gocron"
)

func TxListener(repo repository.Repository, db repository.DB) error {

	err := gocron.Every(10).Second().Do(ListenTx, repo, db)
	return err
}

func FinancialListener(repo repository.Repository, db repository.DB) error {
	err := gocron.Every(30).Second().Do(ExecuteFiancialTx, repo, db)
	// err := ExecuteFiancialTx(repo, db)
	return err
}

func ValidationListener(repo repository.Repository, db repository.DB) error {

	err := gocron.Every(5).Second().Do(ValidateTx, repo, db)

	return err
}

func CheckValidationListener(repo repository.Repository, db repository.DB) error {
	err := gocron.Every(30).Second().Do(CheckValidation, repo, db)

	return err
}

func ExecuteFiancialTx(repo repository.Repository, db repository.DB) error {

	pendingTx, err := repo.FindTx(service.AddTxCollection, "pending")
	if err != nil {
		return err
	}
	for _, tx := range pendingTx {
		go GoRoutineTest(repo, tx, db)
	}
	return nil

}
func GoRoutineTest(repo repository.Repository, tx domain.TransactionDetails, db repository.DB) {
	success, PaiementToken, err := service.VerifTransaction(repo, tx, db)
	if err != nil {
		_, err = repo.UpdateDB(tx, success, service.AddTxCollection)
	}

	UpdateFinancialdetails(repo, success, tx, PaiementToken)

}
func UpdateFinancialdetails(repo repository.Repository, success bool, tx domain.TransactionDetails, token string) error {

	_, err := repo.UpdateDB(tx, success, service.AddTxCollection)
	if err != nil {
		return err
	}
	err = repo.UpdatePaiement(tx.Tx.Value.Msg[0].Value.Participator, token, success)
	return err
}

func ValidateTx(repo repository.Repository, db repository.DB) error {
	Txs, err := repo.FindTx(service.AddTxCollection, "true", "false")

	if err != nil {
		return err
	}
	cdc := utils.NewCodec()
	for _, tx := range Txs {
		out, err := service.ConfirmTx(cdc, tx)
		if err != nil {
			return err
		}
		txReceipt := domain.TxReceipt{}
		json.Unmarshal(out, &txReceipt)
		if txReceipt.Code != 0 {
			return errors.New("An error while executing the tx in the blockchain")
		}
		_, err = repo.UpdateValidationDB(tx, true, service.AddTxCollection, txReceipt.Txhash)
		if err != nil {
			return err
		}
	}
	return nil

}

func CheckValidation(repo repository.Repository, db repository.DB) error {
	baseURL := utils.GetRestUrl()
	err := service.GetValidConfirmations(repo, baseURL)
	if err != nil {
		return err
	}
	return nil
}

func ListenTx(repo repository.Repository, db repository.DB) error {
	baseURL := utils.GetRestUrl()
	err := service.GetTxFromBlockchain(repo, baseURL, db)
	if err != nil {
		return err
	}

	return nil
}

func initAdminCreds(repo repository.Repository) error {

	paiement_account := os.Getenv("PAIEMENT_ACCOUNT")
	if paiement_account == "" {
		return errors.New("the paiemet account was not provided")
	}
	authorization_token := os.Getenv("AUTHORIZATION_TOKEN")
	if authorization_token == "" {
		return errors.New("the authorization token was  not provided")
	}
	admincred := domain.AdminCreds{
		PaymeeAccount: paiement_account,
		AuthToken:     authorization_token,
	}
	_, err := repo.RegisterAdminCreds(&admincred)
	if err != nil {
		return err
	}
	return nil

}

func New() error {

	db, err := repository.NewDB()
	if err != nil {
		return err
	}
	ctx := context.Background()
	repo := repository.New(ctx, db)
	initAdminCreds(repo)

	TxListener(repo, db)
	FinancialListener(repo, db)
	ValidationListener(repo, db)
	CheckValidationListener(repo, db)
	<-gocron.Start()
	return nil
}
