package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	if mailchimpSecret == "" {
		panic("MAILCHIMP_SECRET not set")
	}

	if googleFormActionUrl == "" {
		panic("GOOGLE_FORM_ACTION_URL not set")
	}

	if googleFormEmailField == "" {
		panic("GOOGLE_FORM_EMAIL_FIELD not set")
	}
	router = mux.NewRouter()

	router.HandleFunc("/{secret}", handleMailchimpWebhookPing).Methods("GET")
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

var mailchimpSecret string = os.Getenv("MAILCHIMP_SECRET")
var googleFormActionUrl string = os.Getenv("GOOGLE_FORM_ACTION_URL")
var googleFormEmailField string = os.Getenv("GOOGLE_FORM_EMAIL_FIELD")
