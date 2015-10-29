package variables

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func TestBenchmarkBcrypt(t *testing.T) {
	fmt.Println("Lowest cost is", BcryptLowest, "computed is", BcryptCost)
	t0 := time.Now()
	bcrypt.GenerateFromPassword([]byte("password"), BcryptCost)
	t1 := time.Now()
	duration := t1.Sub(t0)
	fmt.Println("Time expected", BcryptLowestTime)
	fmt.Println("Time real    ", duration)
	if duration < BcryptLowestTime {
		panic(errors.New("Time expected was lower than time real"))
	}
}
