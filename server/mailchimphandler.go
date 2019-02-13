package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// handleMailchimpWebhookPing returns HTTP 200 OK if the given secret is verified. Mailchimp
// makes an initial testing request to check that the webhook URL is alive.
func handleMailchimpWebhookPing(w http.ResponseWriter, r *http.Request) {
	if validateSecret(w, r) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}

func handleMailchimpWebhook(w http.ResponseWriter, r *http.Request) {
	if !validateSecret(w, r) {
		return
	}

	err := r.ParseForm()

	if err != nil {
		log.Printf("error parsing form: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := r.PostFormValue("data[email]")
	if email == "" {
		err := fmt.Errorf("missing form field data[email]")
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	webhookType := r.PostFormValue("type")
	if webhookType != "subscribe" {
		err := fmt.Errorf("form field type should be `subscribe`, got `%s`", webhookType)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := postToGoogleForm(email); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}

func validateSecret(w http.ResponseWriter, r *http.Request) bool {
	secret := mux.Vars(r)["secret"]
	if secret != mailchimpSecret {
		log.Printf("got bad secret `%s`, expected `%s`", secret, mailchimpSecret)
		http.Error(w, "bad secret", http.StatusForbidden)
		return false
	}

	return true
}

func postToGoogleForm(email string) error {
	log.Printf("posting to google form: %s", email)

	response, err := http.PostForm(googleFormActionUrl, url.Values{
		googleFormEmailField: {email},
	})
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("google forms URL returned HTTP %d", response.StatusCode)
	}
	return nil
}
