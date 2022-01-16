package rest

import (
	"net/http"

	"github.com/cagnotteApp/x/cagnotte/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type addCagnotteRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	User    string       `json:"user"`
	Name    string       `json:"name"`
	Amount  int          `json:"price"`
}

func addCagnotteHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addCagnotteRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		addr, err := sdk.AccAddressFromBech32(req.User)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		coins := req.Amount
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgAddCagnotte(req.Name, coins, addr)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type createCagnotteRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name    string       `json:"name"`
	Creator string       `json:"owner"`
}

func createCagnotteHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createCagnotteRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		addr, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgCreateCagnotte(req.Name, addr)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type closeCagnotteRequest struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Executer string       `json:"executer"`
	Name     string       `json:"name"`
}

func closeCagnotteHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req closeCagnotteRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		addr, err := sdk.AccAddressFromBech32(req.Executer)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgCloseCagnotte(req.Name, addr)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type confirmTxRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Name         string       `json:"cagnotte_name"`
	Amount       int          `json:"amount"`
	Result       bool         `json:"result"`
	Participator string       `json:"participator"`
}

func confirmTxHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req confirmTxRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		succes := req.Result
		userAddr, err := sdk.AccAddressFromBech32(req.Participator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		coins := req.Amount
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgConfirmPendingTx(req.Name, coins, addr, succes, userAddr)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
