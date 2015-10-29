package recaptcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

// XURLRequest makes a x-www-form-urlencoded request to a url
func XURLRequest(urlPath string, data map[string]string, result interface{}) error {
	sendData := url.Values{}
	for index, item := range data {
		sendData.Add(index, item)
	}

	client := http.Client{}
	// FIXME: Need to make sure NewBufferString is the best way to do this
	req, err := http.NewRequest(
		"POST",
		urlPath,
		bytes.NewBufferString(sendData.Encode()),
	)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("XURLRequest error", err)
		return err
	}
	if result == nil {
		return nil
	} else if len(resp.Header["Content-Length"]) > 1 &&
		resp.Header["Content-Length"][0] == "0" {
		return nil
	} else if len(resp.Header["Content-Type"]) > 1 &&
		resp.Header["Content-Type"][0] != "application/json" {
		return errors.New("Response was not json: " + resp.Header["Content-Type"][0])
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}
