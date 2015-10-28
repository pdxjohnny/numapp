package variables

import "os"

const (
	// EnvRecaptchaSecret is the environment variable that
	// stores the reCAPTCHA secret
	EnvRecaptchaSecret = "REC_SECRET"
)

var (
	// RecaptchaSecret is the variable that stores the reCAPTCHA secret
	RecaptchaSecret string
)

func init() {
	RecaptchaSecret = os.Getenv(EnvRecaptchaSecret)
}
