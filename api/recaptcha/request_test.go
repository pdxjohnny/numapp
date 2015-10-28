package recaptcha

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestXURLRequest(t *testing.T) {
	var data = map[string]string{
		"test": "value",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		sendBack := make(map[string]string)
		recvBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		recvString := string(recvBytes)
		received, err := url.ParseQuery(recvString)
		if err != nil {
			panic(err)
		}
		for index, item := range received {
			sendBack[index] = item[0]
		}
		sendBackString, err := json.Marshal(&sendBack)
		if err != nil {
			panic(err)
		}
		w.Write(sendBackString)
	}))
	defer ts.Close()
	var result map[string]string
	err := XURLRequest(ts.URL, data, &result)
	if err != nil {
		panic(err)
	}
	if result["test"] != data["test"] {
		panic(errors.New("Data sent was not the same as received"))
	}
}
