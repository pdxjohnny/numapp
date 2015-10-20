package put

import (
	"errors"

	"github.com/pdxjohnny/numapp/db/shared"
	"github.com/pdxjohnny/numapp/variables"
	"gopkg.in/mgo.v2/bson"
)

// Put tries to insert then tries to save
func Put(collectionName string, doc interface{}) error {
	err := Insert(collectionName, doc)
	if err == nil {
		return nil
	}
	return Update(collectionName, doc)
}

// Insert creates a document
func Insert(collectionName string, doc interface{}) error {
	collection := shared.MongoConn.DB(variables.DBName).C(collectionName)
	err := collection.Insert(doc)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a document
func Update(collectionName string, doc interface{}) error {
	var asMap map[string]interface{}
	switch value := doc.(type) {
	case *map[string]interface{}:
		asMap = *(value)
	case map[string]interface{}:
		asMap = value
	default:
		return errors.New("Must be a map")
	}
	_, ok := asMap["_id"]
	if !ok {
		return errors.New("Doc needs to have _id to be saved")
	}
	findDoc := bson.M{"_id": asMap["_id"]}

	collection := shared.MongoConn.DB(variables.DBName).C(collectionName)
	err := collection.Update(findDoc, doc)
	if err != nil {
		return err
	}

	return nil
}
