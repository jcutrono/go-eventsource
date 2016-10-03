package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/account", newAccount).Methods("POST")
}

type CreateAccount struct {
	Name    string
	Balance int
}

type AccountCreated struct {
	Name    string
	Balance int
}

func newAccount(resp http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var account CreateAccount
	decoder.Decode(&account)

	fmt.Println(account)

	if account.Name == "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	if account.Balance < 0 {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	accountStr, _ := json.Marshal(AccountCreated{
		Name:    account.Name,
		Balance: account.Balance,
	})

	evt := BankEvent{
		UserId:    0,
		AccountId: 0,
		Data:      string(accountStr),
	}

	// Persist First
	write(evt)

	// Publish Second
	evtStr, _ := json.Marshal(evt)
	PublishEvent(string(evtStr))
}
