package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cagnotteApp/x/cagnotte"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/tendermint/go-amino"
)

func Request(method string, url string, body []byte, headers map[string]string) ([]byte, error) {
	requestBody, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			requestBody.Header.Set(k, v)
		}
	}

	requestBody.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(requestBody)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetRestUrl() string {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:1317"
	}
	return baseURL
}

func NewCodec() *amino.Codec {
	ModuleBasics := module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		cagnotte.AppModuleBasic{},
	)

	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
