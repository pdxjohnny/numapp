package login

import (
	"errors"

	"github.com/jmcvetta/randutil"
	"golang.org/x/crypto/bcrypt"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/variables"
)

// Login checks a users password for a match an returns a token if it is
func Login(loginDoc map[string]string) (string, error) {
	id, ok := loginDoc["username"]
	if ok != true {
		return "", errors.New("Need a username to login")
	}
	password, ok := loginDoc["password"]
	if ok != true {
		return "", errors.New("Need a password to login")
	}
	// Get the users account
	doc, err := api.GetAccount(variables.ServiceDBURL, id)
	if err != nil || doc != nil {
		return "", errors.New("Could not find username")
	}
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return "", err
	}
	// Passwords match so create an auth token
	loginToken, err := Token(id)
	if err != nil {
		return "", err
	}
	return loginToken, nil
}

// Token generates an authentication token for the user
func Token(id string) (string, error) {
	token, err := randutil.AlphaString(variables.TokenLength)
	if err != nil {
		return "", err
	}
	return token, nil
}
