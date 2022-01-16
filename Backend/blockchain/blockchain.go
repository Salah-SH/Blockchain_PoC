package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/cagnotteApp/Backend/domain"
	"github.com/cagnotteApp/Backend/utils"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/go-amino"
)

func CreateInitialTx(cdc *amino.Codec, addrTo string, baseURL string) (types.StdTx, error) {
	var unsignedtx types.StdTx
	admin, err := FindAdminAddr()
	if err != nil {
		return unsignedtx, err
	}

	txModel := domain.InitialiseSendTx(admin, "namechain", "2", "1", "10")

	byteValue, err := json.Marshal(txModel)
	if err != nil {
		return unsignedtx, err

	}
	bankUrl := baseURL + fmt.Sprintf("/bank/accounts/%s/transfers", addrTo)

	response, err := utils.Request("POST", bankUrl, byteValue, nil)
	if err != nil {
		return unsignedtx, err

	}
	err = cdc.UnmarshalJSON(response, &unsignedtx)
	if err != nil {
		return unsignedtx, err
	}
	return unsignedtx, nil
}

func CreateConfirmTx(cdc *amino.Codec, baseURL string, tx domain.TransactionDetails, result bool) (types.StdTx, error) {
	var unsignedtx types.StdTx
	admin, err := FindAdminAddr()
	if err != nil {
		return unsignedtx, err
	}

	txModel := domain.InitialiseConfirmTx(admin, "namechain", "2", "1", tx.Tx.Value.Msg[0].Value.Bid, tx.Tx.Value.Msg[0].Value.Name, tx.Tx.Value.Msg[0].Value.Participator, result)
	byteValue, err := json.Marshal(txModel)
	if err != nil {
		return unsignedtx, err

	}
	Url := baseURL + "/cagnotte/confirmtx"
	response, err := utils.Request("POST", Url, byteValue, nil)
	if err != nil {
		return unsignedtx, err

	}
	err = cdc.UnmarshalJSON(response, &unsignedtx)

	if err != nil {
		return unsignedtx, err
	}

	return unsignedtx, nil
}

func findSeqAndAccNum(baseUrl string) (int64, int64, error) {

	admin, err := FindAdminAddr()
	url := fmt.Sprintf("%s/auth/accounts/%s", baseUrl, admin)

	response, err := utils.Request("GET", url, nil, nil)
	if err != nil {
		return 0, 0, err
	}
	adminAccount := domain.AccountStruct{}
	err = json.Unmarshal(response, &adminAccount)
	if err != nil {
		return 0, 0, err
	}
	sequence, _ := strconv.ParseInt(adminAccount.Result.Value.Sequence, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	accountNumber, _ := strconv.ParseInt(adminAccount.Result.Value.AccountNumber, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return sequence, accountNumber, nil
}

func SignTx(stdTx types.StdTx, baseUrl string) (types.StdTx, error) {

	var signedStdTx types.StdTx
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		return signedStdTx, errors.New("keys directory not found")
	}
	sequence, accountNum, err := findSeqAndAccNum(baseUrl)
	if err != nil {
		return signedStdTx, err
	}
	kb, err := keys.NewKeyring("cagnotte", "test", rootDir, nil)
	if err != nil {
		return signedStdTx, err
	}
	sigBytes, pubkey, err := kb.Sign("admin", "", types.StdSignMsg{
		ChainID:       "namechain",
		AccountNumber: uint64(accountNum),
		Fee:           types.NewStdFee(200000, nil),
		Sequence:      uint64(sequence),
		Msgs:          stdTx.GetMsgs(),
	}.Bytes())

	if err != nil {
		return signedStdTx, err
	}

	stdSignature := types.StdSignature{
		PubKey:    pubkey,
		Signature: sigBytes,
	}

	sigs := stdTx.Signatures
	sigs = []types.StdSignature{stdSignature}
	signedStdTx = types.NewStdTx(stdTx.GetMsgs(), stdTx.Fee, sigs, stdTx.GetMemo())
	return signedStdTx, nil

}

func FindAdminAddr() (string, error) {
	admin := os.Getenv("ADMIN_ADDR")
	if admin == "" {
		return "", errors.New("the admin address is not provided")
	}
	return admin, nil

}
