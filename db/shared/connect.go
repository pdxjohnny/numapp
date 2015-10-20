package shared

import (
	"log"
	"os"

	"github.com/pdxjohnny/numapp/variables"

	"gopkg.in/mgo.v2"
)

// MongoConn is the conenction to the mongo server
var MongoConn *mgo.Session

// init establishes a conenction with mongodb
func init() {
	uri := os.Getenv(variables.DBAddress)
	if uri == "" {
		MongoConn = nil
		log.Println("No mongo server to connect to")
		return
	}

	MongoConn, err := mgo.Dial(uri)
	if err != nil {
		MongoConn = nil
		log.Println("Could not connect to mongo server")
		return
	}

	// Optional. Switch the MongoConn to a monotonic behavior.
	MongoConn.SetMode(mgo.Monotonic, true)
	MongoConn.SetSafe(&mgo.Safe{})

	return
}
