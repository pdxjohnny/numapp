package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/db/put"
)

// PostDoc uses put to retrive a document
func PostDoc(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")
	var doc map[string]interface{}
	err := r.DecodeJsonPayload(&doc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = put.Put(collection, doc)
	if err != nil {
		log.Println(err)
		rest.Error(w, "Could not save", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(nil)
}
