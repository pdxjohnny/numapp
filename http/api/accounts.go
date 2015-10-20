package api

import (
	"log"
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
	}
}

// PostAccount uses get to retrive a document
func PostAccount(w rest.ResponseWriter, r *rest.Request) {
	var doc map[string]interface{}
	err := r.DecodeJsonPayload(&doc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.SaveAccount(variables.ServiceDBURL, doc)
	if err != nil {
		log.Println(err)
		rest.Error(w, "Could not save", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(nil)
}
