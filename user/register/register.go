package register

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/api/recaptcha"
	"github.com/pdxjohnny/numapp/variables"
)

// Register takes a doc and attempts to create a new user
func Register(registerDoc map[string]interface{}) error {
	id, ok := registerDoc["username"].(string)
	if ok != true {
		return errors.New("Need username to register")
	}
	passwordString, ok := registerDoc["password"].(string)
	if ok != true {
		return errors.New("Need password to register")
	}
	password := []byte(passwordString)
	reCAPTCHA, ok := registerDoc["g-recaptcha-response"].(string)
	if ok != true {
		return errors.New("Need g-recaptcha-response to register")
	}
	doc, err := api.GetUser(variables.ServiceDBURL, id)
	if err != nil || doc != nil {
		return errors.New("Username is already taken")
	}
	// Verify with google reCAPTCHA
	err = recaptcha.Verify(variables.RecaptchaSecret, reCAPTCHA)
	if err != nil {
		return errors.New("reCAPTCHA invalid")
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		variables.BcryptCost,
	)
	if err != nil {
		return err
	}
	registerDoc["password"] = hashedPassword

	_, err = api.SaveUser(variables.ServiceDBURL, id, registerDoc)
	if err != nil {
		return err
	}
	return nil
}
