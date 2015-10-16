package put

import (
	"errors"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// Put saves a document
func Put(doc interface{}) error {
	uri := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	log.Println("Connecting to", uri)
	if uri == "" {
		return errors.New("No mongo server to connect to")
	}

	session, err := mgo.Dial(uri)
	if err != nil {
		return err
	}
	log.Println("Connected to", uri)
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	session.SetSafe(&mgo.Safe{})
	collection := session.DB("numapp").C("numbers")
	err = collection.Insert(doc)
	if err != nil {
		return err
	}

	return nil
}
