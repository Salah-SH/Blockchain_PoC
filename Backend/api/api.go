package api

import (
	"encoding/json"
	"net/http"

	"github.com/cagnotteApp/Backend/domain"
	"github.com/cagnotteApp/Backend/repository"
	"github.com/cagnotteApp/Backend/service"
	"github.com/cagnotteApp/Backend/utils"
	"github.com/gorilla/mux"
	"github.com/tendermint/go-amino"
)

func createUser(w http.ResponseWriter, r *http.Request, cdc *amino.Codec, db repository.DB) {
	ctx := r.Context()
	user := domain.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errorResponse(http.StatusBadRequest, err.Error()))

		return
	}
	w.Header().Set("Content-Type", "application/json")

	repo := repository.New(ctx, db)
	resp, err := service.RegisterUser(cdc, repo, user)

	if err != nil {

		writeResponse(w, http.StatusBadRequest, errorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	writeResponse(w, http.StatusOK, ApiResponse{
		Success:  true,
		Response: resp,
	})

}

func getPaiements(w http.ResponseWriter, r *http.Request, cdc *amino.Codec, db repository.DB) {
	ctx := r.Context()
	vars := mux.Vars(r)
	accAddr := vars["address"]

	repo := repository.New(ctx, db)

	resp, err := service.GetPaiements(repo, accAddr)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errorResponse(http.StatusBadRequest, err.Error()))

	}
	writeResponse(w, http.StatusOK, ApiResponse{
		Success:  true,
		Response: resp,
	})
}

func New() error {
	db, err := repository.NewDB()
	if err != nil {
		return err
	}
	cdc := utils.NewCodec()
	HandleRequests(cdc, db)

	return nil
}

type ApiResponse struct {
	Success  bool        `json:"success"`
	Error    *APIError   `json:"error,omitempty"`
	Response interface{} `json:"response,omitempty"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

func writeResponse(w http.ResponseWriter, statusCode int, response ApiResponse) {
	encoder := json.NewEncoder(w)
	//	w.WriteHeader(statusCode)
	encoder.Encode(response)
}

func errorResponse(code int, message string) ApiResponse {
	return ApiResponse{Error: &APIError{Code: code, Message: message}}
}
