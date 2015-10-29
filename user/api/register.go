package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// PostRegister registers a new user
func PostRegister(w rest.ResponseWriter, r *rest.Request) {
	var registerReq map[string]interface{}
	err := r.DecodeJsonPayload(&registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	// if doc == nil {
	// 	w.(http.ResponseWriter).Write(variables.BlankResponse)
	// 	return
	// }
	// w.WriteJson(doc)
}
