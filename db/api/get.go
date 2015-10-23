package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/db/get"
	"github.com/pdxjohnny/numapp/variables"
)

// GetDoc uses get to retrive a document
func GetDoc(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	collection := r.PathParam("collection")
	doc, err := get.Get(collection, id)
	if err != nil {
		rest.Error(w, "Not Found", 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(variables.BlankResponse)
		return
	}
	w.WriteJson(doc)
}
