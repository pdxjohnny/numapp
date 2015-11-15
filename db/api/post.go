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
	id := r.PathParam("id")
	collection := r.PathParam("collection")
	// Make sure that the user owns what they are trying to save
	_, ok := r.Env["JWT_PAYLOAD"].(map[string]interface{})["backend"]
	if !ok && r.Env["REMOTE_USER"] != id {
		rest.Error(w, "Cannot modify another users data", http.StatusInternalServerError)
		return
	}
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
