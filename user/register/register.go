package register

import (
	"errors"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/api/recaptcha"
	"github.com/pdxjohnny/numapp/variables"
)

// Register takes a doc and attempts to create a new user
func Register(registerDoc map[string]interface{}) error {
	// Make sure we have all the properties we need
	err := checkPropsExist(registerDoc)
	if err != nil {
		return err
	}
	// Get the properties we need
	id := registerDoc["username"].(string)
	email := registerDoc["email"].(string)
	password := registerDoc["password"].(string)
	reCAPTCHA := registerDoc["g-recaptcha-response"].(string)
	// Make sure the user does not exist already
	doc, err := api.GetUser(variables.ServiceDBURL, id)
	if doc != nil {
		return errors.New("Username is already taken")
	}
	// Verify with google reCAPTCHA
	err = recaptcha.Verify(variables.RecaptchaSecret, reCAPTCHA)
	if err != nil {
		return errors.New(err.Error())
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		variables.BcryptCost,
	)
	if err != nil {
		return err
	}

	// Take only the parts we care about and save them
	saveDoc := map[string]interface{}{
		"id":       id,
		"email":    email,
		"password": hashedPassword,
	}

	_, err = api.SaveUser(variables.ServiceDBURL, id, saveDoc)
	if err != nil {
		return err
	}
	log.Println("User registered", saveDoc)
	return nil
}

// checkPropsExist makes sure all the properties that need to exist do
// before attempting to register
func checkPropsExist(registerDoc map[string]interface{}) error {
	// Make sure the user provides a username
	id, ok := registerDoc["username"].(string)
	if ok != true {
		return errors.New("Need username to register")
	} else if len(id) < variables.ShortestUsername {
		// FIXME: Make concatination constant
		return errors.New("Username must be at least " + strconv.Itoa(variables.ShortestUsername))
	}
	// Make sure the user provides a password
	password, ok := registerDoc["password"].(string)
	if ok != true {
		return errors.New("Need password to register")
	} else if len(password) < variables.ShortestPassword {
		// FIXME: Make concatination constant
		return errors.New("Password must be at least " + strconv.Itoa(variables.ShortestPassword))
	}
	// Make sure the passwords match
	confirmPassword, ok := registerDoc["confirm_password"].(string)
	if ok != true {
		return errors.New("Need confirm_password to register")
	} else if password != confirmPassword {
		return errors.New("Passwords do not match")
	}
	// Make sure the user provides an email
	_, ok = registerDoc["email"].(string)
	if ok != true {
		return errors.New("Need email to register")
	}
	// Make sure the user provides a recaptcha
	_, ok = registerDoc["g-recaptcha-response"].(string)
	if ok != true {
		return errors.New("Need g-recaptcha-response to register")
	}
	return nil
}
