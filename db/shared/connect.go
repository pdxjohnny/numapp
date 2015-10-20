package shared

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// MongoConn is the conenction to the mongo server
var MongoConn *mgo.Session

// init establishes a conenction with mongodb
func init() {
	uri := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	log.Println("Connecting to mongo server", uri)
	if uri == "" {
		log.Println("No mongo server to connect to")
		return
	}

	MongoConn, err := mgo.Dial(uri)
	if err != nil {
		log.Println("Could not connect to mongo server")
		return
	}
	log.Println("Connected to mongo server", uri)

	// Optional. Switch the MongoConn to a monotonic behavior.
	MongoConn.SetMode(mgo.Monotonic, true)
	MongoConn.SetSafe(&mgo.Safe{})

	return
}
