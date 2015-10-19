package api

import (
	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/get"
)

// GetDoc uses get to retrive a document
func GetDoc(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	doc, err := get.Get(id)
	if err != nil {
		rest.Error(w, "Not Found", 404)
	}
	w.WriteJson(doc)
}
