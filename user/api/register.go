package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pdxjohnny/numapp/user/register"
	"github.com/pdxjohnny/numapp/variables"
)

// PostRegister registers a new user
func PostRegister(w rest.ResponseWriter, r *rest.Request) {
	log.Println("Got register req")
	var registerReq map[string]interface{}
	err := r.DecodeJsonPayload(&registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Registering")
	err = register.Register(registerReq)
	if err != nil {
		log.Println("Error registering", err.Error())
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if doc == nil {
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(variables.BlankResponse)
	// 	return
	// }
	// w.WriteJson(doc)
}
