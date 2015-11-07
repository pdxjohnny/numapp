package api

import (
	"strings"

	"github.com/pdxjohnny/numapp/variables"
)

// LoginUser logs in a user
func LoginUser(host, token string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathLoginUser
	return GenericRequest(host, path, token, doc)
}

// RefreshUser gets a new toke for the user
func RefreshUser(host, token string) (*map[string]interface{}, error) {
	path := variables.APIPathRefreshUser
	return GenericRequest(host, path, token, nil)
}

// RegisterUser registers a user
func RegisterUser(host, token string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathRegisterUser
	return GenericRequest(host, path, token, doc)
}

// GetUser retrives a user
func GetUser(host, token, id string) (*map[string]interface{}, error) {
	path := variables.APIPathUser
	path = strings.Replace(path, ":id", id, 1)
	return GenericRequest(host, path, token, nil)
}

// SaveUser saves a user
func SaveUser(host, token, id string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathUser
	path = strings.Replace(path, ":id", id, 1)
	doc["_id"] = id
	return GenericRequest(host, path, token, doc)
}
