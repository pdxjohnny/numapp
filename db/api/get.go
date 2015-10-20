package api

import (
	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/db/get"
)

// GetDoc uses get to retrive a document
func GetDoc(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	collection := r.PathParam("collection")
	doc, err := get.Get(collection, id)
	if err != nil {
		rest.Error(w, "Not Found", 404)
	}
	if doc != nil {
		w.WriteJson(doc)
	}
}
