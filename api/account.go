package api

import (
	"errors"
	"strings"

	"github.com/pdxjohnny/numapp/variables"
)

// GetAccount calls database handler and retrives an account
func GetAccount(host string, id string) (*map[string]interface{}, error) {
	var result map[string]interface{}
	if host == "" {
		return nil, errors.New("Host is blank")
	}
	host += variables.APIPathAccount
	url := strings.Replace(host, ":id", id, 1)
	resp, err := RESTRequest(url, nil, &result)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, errors.New("Status: " + resp.Status)
	}
	return &result, nil
}

// SaveAccount calls database handler and saves an account
func SaveAccount(url string, doc map[string]interface{}) error {
	return nil
}
