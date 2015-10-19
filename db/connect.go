package db

import (
	"errors"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// Connect establishes a conenction with mongodb
func Connect() (*mgo.Session, error) {
	uri := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	log.Println("Connecting to", uri)
	if uri == "" {
		return nil, errors.New("No mongo server to connect to")
	}

	session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to", uri)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})

	return session, nil
}
