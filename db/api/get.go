package api

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/db/get"
)

// GetDoc uses get to retrive a document
func GetDoc(w rest.ResponseWriter, r *rest.Request) {
	log.Println("Geting id param")
	id := r.PathParam("id")
	log.Println("Geting collection param")
	collection := r.PathParam("collection")
	log.Println("Geting", id, "from", collection)
	doc, err := get.Get(collection, id)
	if err != nil {
		rest.Error(w, "Not Found", 404)
	}
	if doc != nil {
		w.WriteJson(doc)
	}
}
