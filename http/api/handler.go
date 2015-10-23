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
		// For accounts, looking up and updating
		rest.Get(variables.APIPathAccountServer, GetAccount),
		rest.Post(variables.APIPathAccountServer, PostAccount),
		// For user actions such as login
		rest.Post("/user/login/:username", PostUserLogin),
		rest.Post("/user/register/:username", PostUserRegister),
		rest.Get(variables.APIPathUserSettingsServer, GetUserSettings),
		rest.Post(variables.APIPathUserSettingsServer, PostUserSettings),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := api.MakeHandler()
	return &handler
}
