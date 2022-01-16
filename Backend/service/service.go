package service

import (
	"encoding/json"
	"errors"

	"github.com/cagnotteApp/Backend/domain"
	"github.com/cagnotteApp/Backend/repository"
	"github.com/tendermint/go-amino"
)

func RegisterUser(cdc *amino.Codec, repo repository.Repository, user domain.User) (interface{}, error) {
	if client, err := repo.FindUser(user.AccAddr); err != nil || client != nil {

		if err != nil {
			if err.Error() != "user not found in the db" {
				return nil, err
			}

		} else if client != nil {
			return nil, errors.New("User already registered")
		}
	}

	KYC := kycValidation(user)

	if !KYC {
		return nil, errors.New("Cannot verify user information")
	}

	out, err := initTransaction(cdc, user.AccAddr)
	txReceipt := domain.TxReceipt{}
	json.Unmarshal(out, &txReceipt)
	if err != nil {
		return nil, err
	}

	resp, err := repo.SaveUser(&user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetPaiements(repo repository.Repository, address string) (interface{}, error) {
	client, err := repo.FindUser(address)
	if err != nil {
		return nil, err
	}
	activePaiments := []domain.Paiement{}
	for _, value := range client.Paiements {
		if value.Status == "not paid" {
			activePaiments = append(activePaiments, value)
		}
	}
	return activePaiments, nil

}

func RegisterCagnotte(repo repository.Repository, cagnotte *domain.Cagnotte) (*domain.Cagnotte, error) {

	if client, _ := repo.FindUser(cagnotte.Owner); client == nil {
		return nil, errors.New("User not registered")

	}
	cgnt, err := repo.FindCagnotte(cagnotte.Name)
	if cgnt != nil {

		return cgnt, nil
	}
	if err.Error() != "cagnotte not found in db" {
		return nil, err
	}
	cgnt, err = repo.SaveCagnotte(cagnotte)
	return cgnt, err

}
