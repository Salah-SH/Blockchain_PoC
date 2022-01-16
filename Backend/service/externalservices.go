package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cagnotteApp/Backend/domain"
	"github.com/cagnotteApp/Backend/repository"
	"github.com/cagnotteApp/Backend/utils"
)

func kycValidation(user domain.User) bool {
	if user.AccAddr[len(user.AccAddr)-1:] == "a" {

		return false
	}
	return true

}

func createFinancialTx(repo repository.Repository, user domain.User, tx domain.TransactionDetails) (bool, string, error) {
	admin, err := repo.GetAdminCreds()
	if err != nil {
		return false, "", errors.New("The admin creds are not found")
	}
	paiement, err := CreatePaiement(admin, tx.Tx.Value.Msg[0].Value.Bid, tx.Tx.Value.Msg[0].Value.Name)
	if err != nil {
		return false, "", err
	}
	if paiement.Status != true {
		return false, "", err
	}

	err = repo.RegisterPaiement(tx.Tx.Value.Msg[0].Value.Participator, paiement.Data.Token, tx.Tx.Value.Msg[0].Value.Name, tx.Tx.Value.Msg[0].Value.Bid)
	if err != nil {
		return false, "", err
	}
	return true, paiement.Data.Token, nil

}

func VerifPaiementExecution(repo repository.Repository, paiemenToken string) (bool, error) {
	admin, err := repo.GetAdminCreds()
	if err != nil {
		return false, errors.New("The admin creds are not found")
	}
	paiement, err := verifPaymeePaiement(admin, paiemenToken)
	return paiement.Data.PaymentStatus, nil
}

func RegisterTx(repo repository.Repository, txs []domain.TransactionDetails, dbCollection string) (bool, error) {
	for _, tx := range txs {
		_, err := repo.SaveAddTx(&tx, dbCollection)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func RegisterConfirmedTx(repo repository.Repository, txs []domain.TransactionDetails, dbCollection string) (bool, error) {
	for _, tx := range txs {
		_, err := repo.SaveConfirmedTx(&tx, dbCollection)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func CreatePaiement(admin domain.AdminCreds, amount string, cagnotteName string) (domain.CreatePaiementResponse, error) {
	PaiementReceipt := domain.CreatePaiementResponse{}

	headers := make(map[string]string)
	authToken := fmt.Sprintf("Token %s", admin.AuthToken)
	headers["Authorization"] = authToken
	request := domain.CreatePaiement{
		Vender: admin.PaymeeAccount,
		Amount: amount,
		Note:   "Participation in Jackpot :  " + cagnotteName,
	}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return PaiementReceipt, err
	}
	url := "https://sandbox.paymee.tn/api/v1/payments/create"
	response, err := utils.Request("POST", url, requestBody, headers)
	if err != nil {
		return PaiementReceipt, err
	}
	err = json.Unmarshal(response, &PaiementReceipt)
	if err != nil {
		return PaiementReceipt, err
	}
	return PaiementReceipt, nil
}
func verifPaymeePaiement(admin domain.AdminCreds, paiementoken string) (domain.PaymeeResponse, error) {
	paymeeAPI := "https://sandbox.paymee.tn/api/v1/payments"

	paymeeResponse := domain.PaymeeResponse{}
	headers := make(map[string]string)
	authToken := fmt.Sprintf("Token %s", admin.AuthToken)
	headers["Authorization"] = authToken
	url := fmt.Sprintf("%s/%s/check", paymeeAPI, paiementoken)
	response, err := utils.Request("GET", url, nil, headers)
	if err != nil {
		return paymeeResponse, err
	}
	err = json.Unmarshal(response, &paymeeResponse)
	if err != nil {

		return paymeeResponse, err
	}
	return paymeeResponse, nil
}
