package variables

import (
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	// EnvRecaptchaSecret is the environment variable that
	// stores the reCAPTCHA secret
	EnvRecaptchaSecret = "REC_SECRET"
	// BcryptLowest is the lowest cost allowed for bcrypt
	BcryptLowest = 5
	// BcryptLowestTime is the lowest time allowed for the bcrypt cost
	BcryptLowestTime = time.Duration(1000 * time.Millisecond)
)

var (
	// RecaptchaSecret is the variable that stores the reCAPTCHA secret
	RecaptchaSecret string
	// BcryptCost is the cost used for bcrypt
	BcryptCost int
)

func benchmarkBcrypt() int {
	cost := BcryptLowest
	t0 := time.Now()
	bcrypt.GenerateFromPassword([]byte("password"), cost)
	t1 := time.Now()
	duration := t1.Sub(t0)
	for ; duration < BcryptLowestTime; cost++ {
		duration *= 2
	}
	if cost < BcryptLowest {
		return BcryptLowest
	}
	return cost
}

func init() {
	RecaptchaSecret = os.Getenv(EnvRecaptchaSecret)
	BcryptCost = benchmarkBcrypt()
}
