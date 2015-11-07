package variables

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadToken(t *testing.T) {
	os.Setenv(EnvTokenSignKey, "../"+TokenSignKeyDefault)
	os.Setenv(EnvTokenVerifyKey, "../"+TokenVerifyKeyDefault)
	err := LoadTokenKeys()
	if err != nil {
		panic(err)
	}
	fmt.Println("BackendToken", BackendToken)
}
