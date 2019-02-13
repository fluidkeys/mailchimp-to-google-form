package server

import (
	"os"
)

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
}

var mailchimpSecret string = os.Getenv("MAILCHIMP_SECRET")
var googleFormActionUrl string = os.Getenv("GOOGLE_FORM_ACTION_URL")
var googleFormEmailField string = os.Getenv("GOOGLE_FORM_EMAIL_FIELD")
