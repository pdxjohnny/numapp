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
		rest.Error(w, err.Error(), 404)
	}
	if doc != nil {
		w.WriteJson(doc)
	} else {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
	}
}

// PostAccount uses get to retrive a document
func PostAccount(w rest.ResponseWriter, r *rest.Request) {
	var doc map[string]interface{}
	id := r.PathParam("id")
	err := r.DecodeJsonPayload(&doc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	savedDoc, err := api.SaveAccount(variables.ServiceDBURL, id, doc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if savedDoc != nil {
		w.WriteJson(savedDoc)
	} else {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
	}
}
