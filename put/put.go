package put

import (
	"errors"

	"github.com/pdxjohnny/numapp/db"
	"gopkg.in/mgo.v2/bson"
)

// Put tries to insert then tries to save
func Put(doc interface{}) error {
	err := Insert(doc)
	if err == nil {
		return nil
	}
	return Update(doc)
}

// Insert creates a document
func Insert(doc interface{}) error {
	session, err := db.Connect()
	if err != nil {
		return err
	}
	defer session.Close()

	collection := session.DB("numapp").C("numbers")
	err = collection.Insert(doc)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a document
func Update(doc interface{}) error {
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

	session, err := db.Connect()
	if err != nil {
		return err
	}
	defer session.Close()

	collection := session.DB("numapp").C("numbers")
	err = collection.Update(findDoc, doc)
	if err != nil {
		return err
	}

	return nil
}
