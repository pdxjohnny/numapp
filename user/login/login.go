package login

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/variables"
)

// Login checks a users password for a match an returns a token if it is
func Login(id, password string) error {
	// Get the users account
	doc, err := api.GetUser(variables.ServiceDBURL, id)
	if err != nil || doc == nil {
		return errors.New("Could not find username")
	}
	// Comparing the password with the hash
	realPasswordString, ok := (*doc)["password"].(string)
	if ok != true {
		return errors.New("No password for user")
	}
	realPassword := []byte(realPasswordString)
	err = bcrypt.CompareHashAndPassword(realPassword, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
