package query

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/account/{id}", getAccount).Methods("GET")
}

func handleAccountCreated(evt AccountCreated) {

	writeAccount(BankAccount{
		Name:    evt.Name,
		Balance: evt.Balance,
	})
}

func getAccount(resp http.ResponseWriter, req *http.Request) {

	id, _ := mux.Vars(req)["id"]
	account := findAccount(id)

	val, _ := json.Marshal(account)

	resp.Write(val)

}
