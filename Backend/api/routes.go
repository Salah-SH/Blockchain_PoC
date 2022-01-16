package api

import (
	"log"
	"net/http"

	"github.com/cagnotteApp/Backend/repository"
	"github.com/gorilla/mux"
	"github.com/tendermint/go-amino"
)

func HandleRequests(cdc *amino.Codec, db repository.DB) {
	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		createUser(w, r, cdc, db)
	}).Methods("POST")
	Router.HandleFunc("/paiement/{address}", func(w http.ResponseWriter, r *http.Request) {
		getPaiements(w, r, cdc, db)

	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", Router))
}
