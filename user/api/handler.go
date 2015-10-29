package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/variables"
)

// MakeHandler creates the api request handler
func MakeHandler() *http.Handler {
	api := rest.NewApi()
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		rest.Post("/user/login/:id", PostLogin),
		rest.Post("/user/register/:id", PostRegister),
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
