package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// GenericRequest makes a request and wraps it in some error handling
func GenericRequest(host, path, token string, data interface{}) (*map[string]interface{}, error) {
	var result map[string]interface{}
	if host == "" {
		return nil, errors.New("500 No host given")
	}
	host += path
	resp, err := RESTRequest(host, token, data, &result)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		// Return the servers error message as the error if we can
		value, ok := result["Error"]
		if ok != true {
			return nil, errors.New(resp.Status)
		}
		return nil, errors.New(value.(string))
	}
	return &result, nil
}

// RESTRequest makes a rest request to a url and posts data as json
func RESTRequest(url, token string, data interface{}, result interface{}) (*http.Response, error) {
	var request *http.Request
	if data != nil {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		reader := bytes.NewBuffer(jsonBytes)
		req, err := http.NewRequest("POST", url, reader)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		request = req
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		request = req
	}
	request.Header.Set("Accept-Encoding", "application/json")
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Make sure there is data to encode
	if result == nil {
		return resp, nil
	}
	if len(resp.Header["Content-Length"]) > 1 &&
		resp.Header["Content-Length"][0] == "0" {
		return resp, nil
	} else if len(resp.Header["Content-Type"]) > 1 &&
		resp.Header["Content-Type"][0] != "application/json" {
		return resp, errors.New("Response was not json: " + resp.Header["Content-Type"][0])
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
