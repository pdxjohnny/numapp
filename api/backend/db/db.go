package db

import (
	"log"
	"os"

	"github.com/pdxjohnny/numapp/variables"
)

// DBServiceURL is the url of the db service we are contacting
var DBServiceURL string

func init() {
	DBServiceURL = os.Getenv(variables.ServiceDBURL)
	if DBServiceURL == "" {
		log.Println("No db service to connect to")
		return
	}
}
