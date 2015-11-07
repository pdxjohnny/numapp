package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/pdxjohnny/numapp/db/save"
	"github.com/pdxjohnny/numapp/variables"
)

// PostDoc uses save to save a document
func PostDoc(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")
	var doc map[string]interface{}
	err := r.DecodeJsonPayload(&doc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(doc)
	err = save.Save(collection, doc)
	if err != nil {
		log.Println(err)
		rest.Error(w, "Could not save", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(variables.BlankResponse)
}
