package api

import (
	"errors"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/variables"
)

// GetUser returns the user for an id
func GetUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	if r.Env["REMOTE_USER"].(string) != id {
		err := errors.New("Can only access your own user account")
		rest.Error(w, err.Error(), http.StatusUnauthorized)
	}
	doc, err := api.GetUser(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), id)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}

// PostUser uses get to retrive a document
func PostUser(w rest.ResponseWriter, r *rest.Request) {
	var recvDoc map[string]interface{}
	id := r.PathParam("id")
	if r.Env["REMOTE_USER"].(string) != id {
		err := errors.New("Can only save your own user account")
		rest.Error(w, err.Error(), http.StatusUnauthorized)
	}
	err := r.DecodeJsonPayload(&recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc, err := api.SaveUser(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), id, recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}
