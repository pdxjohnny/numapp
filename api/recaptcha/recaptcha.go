package recaptcha

import "errors"

const (
	// VerifyURL is the url to POST to, owned by google
	VerifyURL = "https://www.google.com/recaptcha/api/siteverify"
)

// Verify checks with googles verification server to make sure the user
// got the recaptcha correct
func Verify(secret, userResponse string) error {
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
		return errors.New("reCAPTCHA was invalid")
	}
	return nil
}
