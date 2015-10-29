package api

import (
	"strings"

	"github.com/pdxjohnny/numapp/variables"
)

// LoginUser logs in a user
func LoginUser(host string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathLoginUser
	return GenericRequest(host, path, doc)
}

// RegisterUser registers a user
func RegisterUser(host string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathRegisterUser
	return GenericRequest(host, path, doc)
}

// GetUser retrives a user
func GetUser(host, id string) (*map[string]interface{}, error) {
	path := variables.APIPathUser
	path = strings.Replace(path, ":id", id, 1)
	return GenericRequest(host, path, nil)
}

// SaveUser saves a user
func SaveUser(host, id string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathUser
	path = strings.Replace(path, ":id", id, 1)
	doc["_id"] = id
	return GenericRequest(host, path, doc)
}
