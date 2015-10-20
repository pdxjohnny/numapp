package get

import (
	"errors"

	"github.com/pdxjohnny/numapp/db/shared"
	"github.com/pdxjohnny/numapp/variables"

	"gopkg.in/mgo.v2/bson"
)

// Get reteives a record given the id
func Get(collectionName, id string) (*map[string]interface{}, error) {
	connection := shared.Connection()
	if connection == nil {
		return nil, errors.New("No connection to database server")
	}
	var result map[string]interface{}
	collection := connection.DB(variables.DBName).C(collectionName)
	err := collection.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
