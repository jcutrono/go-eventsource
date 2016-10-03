package query

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/account/{id}", getAccount).Methods("GET")
}

func getAccount(resp http.ResponseWriter, req *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	account := findAccount(id)

	val, _ := json.Marshal(account)

	resp.Write(val)

}
