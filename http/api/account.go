package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/variables"
)

// GetAccount returns the accounts for an id
func GetAccount(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	doc, err := api.GetAccount(variables.ServiceDBURL, id)
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

// PostAccount uses get to retrive a document
func PostAccount(w rest.ResponseWriter, r *rest.Request) {
	var saveDoc map[string]interface{}
	id := r.PathParam("id")
	err := r.DecodeJsonPayload(&saveDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc, err := api.SaveAccount(variables.ServiceDBURL, id, saveDoc)
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
