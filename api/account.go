package api

import (
	"errors"
	"strings"

	"gopkg.in/jmcvetta/napping.v3"

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
	resp, err := napping.Get(url, nil, &result, nil)
	if err != nil {
		return nil, err
	} else if resp.Status() != 200 {
		return nil, errors.New("Status: " + string(resp.Status()))
	}
	return &result, nil
}

// SaveAccount calls database handler and saves an account
func SaveAccount(url string, doc map[string]interface{}) error {
	return nil
}
