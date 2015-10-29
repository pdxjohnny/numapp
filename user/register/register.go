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
	log.Println("User registered", registerDoc)
	return nil
}

// checkPropsExist makes sure all the properties that need to exist do
// before attempting to register
func checkPropsExist(registerDoc map[string]interface{}) error {
	id, ok := registerDoc["username"].(string)
	if ok != true {
		return errors.New("Need username to register")
	} else if len(id) < variables.ShortestUsername {
		// FIXME: Make concatination constant
		return errors.New("Username must be at least " + strconv.Itoa(variables.ShortestUsername))
	}
	password, ok := registerDoc["password"].(string)
	if ok != true {
		return errors.New("Need password to register")
	} else if len(password) < variables.ShortestPassword {
		// FIXME: Make concatination constant
		return errors.New("Password must be at least " + strconv.Itoa(variables.ShortestPassword))
	}
	_, ok = registerDoc["g-recaptcha-response"].(string)
	if ok != true {
		return errors.New("Need g-recaptcha-response to register")
	}
	return nil
}
