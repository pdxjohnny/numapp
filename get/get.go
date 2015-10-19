package get

import (
	"github.com/pdxjohnny/numapp/db"

	"gopkg.in/mgo.v2/bson"
)

// Get reteives a record given the id
func Get(id string) (interface{}, error) {
	session, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var result interface{}
	c := session.DB("numapp").C("numbers")
	err = c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
