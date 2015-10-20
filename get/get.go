package get

import (
	"github.com/pdxjohnny/numapp/db"

	"gopkg.in/mgo.v2/bson"
)

// GetAccounts reteives accounts map for an id
func GetAccounts(id string) (interface{}, error) {
	session, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var result interface{}
	c := session.DB("numapp").C("accounts")
	err = c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}

	return GetWith(id, session)
}

// Get reteives a record given the id
func GetWithSession(id string, *mgo.Session) (interface{}, error) {
	err = c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
