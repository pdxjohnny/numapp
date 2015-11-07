package shared

import (
	"log"

	"github.com/pdxjohnny/numapp/variables"

	"gopkg.in/mgo.v2"
)

// MongoConn is the conenction to the mongo server
var MongoConn *mgo.Session
var connectionAttempted bool

func init() {
	MongoConn = nil
	connectionAttempted = false
}

// Connection makes MongoConn accessable
func Connection() *mgo.Session {
	if MongoConn == nil {
		InitConnection()
	}
	return MongoConn
}

// InitConnection establishes a conenction with mongodb
func InitConnection() {
	connectionAttempted = true
	if variables.DBAddress == "" {
		MongoConn = nil
		log.Println("No mongo server to connect to")
		return
	}

	connection, err := mgo.Dial(variables.DBAddress)
	if err != nil {
		connection = nil
		log.Println("Could not connect to mongo server")
		return
	}

	// Optional. Switch the connection to a monotonic behavior.
	connection.SetMode(mgo.Monotonic, true)
	connection.SetSafe(&mgo.Safe{})

	MongoConn = connection
	return
}
