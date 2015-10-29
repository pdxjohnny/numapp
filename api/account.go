package api

import (
	"strings"

	"github.com/pdxjohnny/numapp/variables"
)

// GetAccount retrives an account
func GetAccount(host, id string) (*map[string]interface{}, error) {
	path := variables.APIPathAccount
	path = strings.Replace(path, ":id", id, 1)
	return GenericRequest(host, path, nil)
}

// SaveAccount saves an account
func SaveAccount(host, id string, doc map[string]interface{}) (*map[string]interface{}, error) {
	path := variables.APIPathAccount
	path = strings.Replace(path, ":id", id, 1)
	doc["_id"] = id
	return GenericRequest(host, path, doc)
}
