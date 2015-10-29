package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/user/login"
	"github.com/pdxjohnny/numapp/user/register"
)

// PostRegister registers a new user
func PostRegister(w rest.ResponseWriter, r *rest.Request) {
	var registerReq map[string]interface{}
	err := r.DecodeJsonPayload(&registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = register.Register(registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Now log them in
	auth, err := login.Login(registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(auth)
}
