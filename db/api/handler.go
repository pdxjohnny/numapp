package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pdxjohnny/go-json-rest-middleware-jwt"

	"github.com/pdxjohnny/numapp/variables"
)

// CreateAuthMiddleware creates the middleware for authtication
func CreateAuthMiddleware() (*jwt.Middleware, error) {
	err := variables.LoadTokenVerifyKey()
	if err != nil {
		log.Println("Error loading TokenVerifyKey:", err)
	}

	authMiddleware := &jwt.Middleware{
		Realm:            "numapp",
		SigningAlgorithm: variables.SigningAlgorithm,
		VerifyKey:        variables.TokenVerifyKey,
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour * 24,
		Authenticator: func(username string, password string) error {
			return errors.New("This message should never be seen")
		},
	}
	return authMiddleware, nil
}

// MakeHandler creates the api request handler
func MakeHandler() *http.Handler {
	api := rest.NewApi()

	authMiddleware, err := CreateAuthMiddleware()
	if err != nil {
		panic(err)
	}

	api.Use(&rest.IfMiddleware{
		// Only authenticate non login or register requests
		Condition: func(request *rest.Request) bool {
			return (request.URL.Path != variables.APIPathLoginUserServer) && (request.URL.Path != variables.APIPathRegisterUserServer)
		},
		IfTrue: authMiddleware,
	})
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		rest.Get("/:collection/:id", GetDoc),
		rest.Post("/:collection/:id", PostDoc),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := api.MakeHandler()
	return &handler
}
