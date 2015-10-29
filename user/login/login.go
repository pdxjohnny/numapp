package login

import (
	"errors"

	"github.com/jmcvetta/randutil"
	"golang.org/x/crypto/bcrypt"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/variables"
)

// AuthInfo holds a token and authentication information
type AuthInfo struct {
	token string
}

// Login checks a users password for a match an returns a token if it is
func Login(loginDoc map[string]interface{}) (*AuthInfo, error) {
	id, ok := loginDoc["username"].(string)
	if ok != true {
		return nil, errors.New("Need username to login")
	}
	password, ok := loginDoc["password"].(string)
	if ok != true {
		return nil, errors.New("Need password to login")
	}
	// Get the users account
	doc, err := api.GetUser(variables.ServiceDBURL, id)
	if err != nil || doc == nil {
		return nil, errors.New("Could not find username")
	}
	// Comparing the password with the hash
	realPasswordString, ok := (*doc)["password"].(string)
	if ok != true {
		return nil, errors.New("No password for user")
	}
	realPassword := []byte(realPasswordString)
	err = bcrypt.CompareHashAndPassword(realPassword, []byte(password))
	if err != nil {
		return nil, err
	}
	// Passwords match so create an auth token
	loginToken, err := Token(id)
	if err != nil {
		return nil, err
	}
	// Create the AuthInfo
	auth := &AuthInfo{
		token: loginToken,
	}
	return auth, nil
}

// Token generates an authentication token for the user
func Token(id string) (string, error) {
	token, err := randutil.AlphaString(variables.TokenLength)
	if err != nil {
		return "", err
	}
	return token, nil
}
