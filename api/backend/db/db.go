package db

import (
	"log"
	"os"
)

// DBServiceURL is the url of the db service we are contacting
var DBServiceURL string

func init() {
	DBServiceURL = os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	log.Println("Connecting to db service", DBServiceURL)
	if DBServiceURL == "" {
		log.Println("No db service to connect to")
	}
}
