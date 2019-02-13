package main

import (
	"github.com/fluidkeys/mailchimp-to-google-form/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		panic(err)
	}
}
