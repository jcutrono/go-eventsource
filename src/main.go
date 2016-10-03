package main

import (
	"command"
	"fmt"
	"net/http"
	"os"
	"query"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	initializeApi(router)

	http.Handle("/api/", router)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}

}

func initializeApi(router *mux.Router) {

	// setup api grouping
	apiRoutes := router.PathPrefix("/api").Subrouter()

	command.Configure(apiRoutes)
	query.Configure(apiRoutes)

	command.ConfigurePublish()
	query.ConfigureSubscribe()

	apiRoutes.Headers("Access-Control-Allow-Origin", "*")
	apiRoutes.Headers("Content-Type", "application/json")
}
