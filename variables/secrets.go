package variables

import (
	"io/ioutil"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	// RecaptchaSecretFile is the file that holds the reCAPTCHA secret key
	RecaptchaSecretFile = "keys/reCAPTCHA"
	// BcryptLowest is the lowest cost allowed for bcrypt
	BcryptLowest = bcrypt.DefaultCost
	// BcryptLowestTime is the lowest time allowed for the bcrypt cost
	BcryptLowestTime = time.Duration(1000 * time.Millisecond)
	// TokenLength is the length of an auth token
	TokenLength = 20
	// ShortestUsername is the shortest allowed length for a username
	ShortestUsername = 1
	// ShortestPassword is the shortest allowed length for a password
	ShortestPassword = 6
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
	RecaptchaSecret = ""
	fileData, err := ioutil.ReadFile(RecaptchaSecretFile)
	if err == nil {
		RecaptchaSecret = strings.TrimSpace(string(fileData))
	}
	BcryptCost = benchmarkBcrypt()
}
