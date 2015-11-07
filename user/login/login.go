package login

import (
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/variables"
)

// Login checks a users password for a match an returns a token if it is
func Login(id, password string) error {
	// Get the users account
	doc, err := api.GetUser(variables.ServiceDBURL, variables.BackendToken, id)
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

// checkPropsExist makes sure all the properties that need to exist do
// before attempting to login
func checkPropsExist(loginDoc map[string]interface{}) error {
	// Make sure the user provides a username
	id, ok := loginDoc["username"].(string)
	if ok != true {
		return errors.New("Need username to login")
	} else if len(id) < variables.ShortestUsername {
		// FIXME: Make concatination constant
		return errors.New("Username must be at least " + strconv.Itoa(variables.ShortestUsername))
	}
	// Make sure the user provides a password
	password, ok := loginDoc["password"].(string)
	if ok != true {
		return errors.New("Need password to login")
	} else if len(password) < variables.ShortestPassword {
		// FIXME: Make concatination constant
		return errors.New("Password must be at least " + strconv.Itoa(variables.ShortestPassword))
	}
	return nil
}
