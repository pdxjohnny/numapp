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
	var registerReq map[string]interface{}
	err := r.DecodeJsonPayload(&registerReq)
	if err != nil {
		log.Println("DecodeJsonPayload", err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = register.Register(registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(variables.BlankResponse)
}
