package server

import (
	"os"
)

func init() {
	if mailchimpSecret == "" {
		panic("MAILCHIMP_SECRET not set")
	}

	if googleFormId == "" {
		panic("GOOGLE_FORM_ID not set")
	}

	if googleFormEmailField == "" {
		panic("GOOGLE_FORM_EMAIL_FIELD not set")
	}
}

var mailchimpSecret string = os.Getenv("MAILCHIMP_SECRET")
var googleFormId string = os.Getenv("GOOGLE_FORM_ID")
var googleFormEmailField string = os.Getenv("GOOGLE_FORM_EMAIL_FIELD")
