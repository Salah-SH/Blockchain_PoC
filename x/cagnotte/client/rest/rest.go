package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers cagnotte-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	//this line is used by starport scaffolding
	r.HandleFunc("/cagnotte/participate", addCagnotteHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cagnotte/{key}", getCagnotteHandler(cliCtx, "cagnotte")).Methods("GET")
	r.HandleFunc("/cagnotte/new", createCagnotteHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cagnotte/close", closeCagnotteHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cagnotte/confirmtx", confirmTxHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cagnotte/txs/{address}", getListCagnotteByUserHandler(cliCtx, "cagnotte")).Methods("GET")

}
