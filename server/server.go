package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()

	router.HandleFunc("/{secret}", handleMailchimpWebhook).Methods("POST")
}

func Serve() error {
	http.Handle("/", router)
	return http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
