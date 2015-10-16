package get

import (
	"errors"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Get reteives a record given the id
func Get(id string) (interface{}, error) {
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
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	var result interface{}
	c := session.DB("numapp").C("numbers")
	err = c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
