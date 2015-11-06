package api

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pdxjohnny/go-json-rest-middleware-jwt"

	"github.com/pdxjohnny/numapp/user/login"
	"github.com/pdxjohnny/numapp/variables"
)

// MakeHandler creates the api request handler
func MakeHandler() *http.Handler {
	api := rest.NewApi()

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	authMiddleware := &jwt.Middleware{
		Realm:            "test zone",
		SigningAlgorithm: "ES256",
		Key:              privateKey,
		VerifyKey:        &privateKey.PublicKey,
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour * 24,
		Authenticator: func(username string, password string) error {
			// Log the user in
			err := login.Login(username, password)
			if err != nil {
				return err
			}
			return nil
		},
	}

	// Setup simple app structure
	api.Use(&rest.IfMiddleware{
		// Only authenticate non /login requests
		Condition: func(request *rest.Request) bool {
			return request.URL.Path != variables.APIPathLoginUserServer
		},
		IfTrue: authMiddleware,
	})
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		rest.Post(variables.APIPathLoginUserServer, PostLogin),
		rest.Post(variables.APIPathRegisterUserServer, PostRegister),
		rest.Get(variables.APIPathUserServer, GetUser),
		rest.Post(variables.APIPathUserServer, PostUser),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := api.MakeHandler()
	return &handler
}
