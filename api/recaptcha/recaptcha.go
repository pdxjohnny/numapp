package recaptcha

import (
	"errors"
	"log"
)

const (
	// VerifyURL is the url to POST to, owned by google
	VerifyURL = "https://www.google.com/recaptcha/api/siteverify"
)

// Verify checks with googles verification server to make sure the user
// got the recaptcha correct
func Verify(secret, userResponse string) error {
	// Make sure secret key is not blank
	if len(secret) < 1 {
		return errors.New("Server Error: No reCAPTCHA secret key")
	}
	data := map[string]string{
		"secret":   secret,
		"response": userResponse,
	}
	var result map[string]interface{}
	err := XURLRequest(VerifyURL, data, &result)
	if err != nil {
		return err
	}
	success, ok := result["success"]
	if !ok {
		return errors.New("Response did not contain \"success\"")
	}
	if success != true {
		log.Println("reCAPTCHA verification failed:", result)
		return errors.New("reCAPTCHA was invalid")
	}
	return nil
}
