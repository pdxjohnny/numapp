package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkSame(shouldBe, is map[string]string, key string) {
	if shouldBe[key] != is[key] {
		panic(errors.New("Expected " + shouldBe[key] + " Got" + is[key]))
	}
}

func TestRESTRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	}))
	defer ts.Close()
	_, err := RESTRequest(ts.URL, nil, nil)
	if err != nil {
		panic(err)
	}
}

func TestRESTRequestData(t *testing.T) {
	var data = map[string]string{
		"test": "value",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var result map[string]string
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&result)
		if err != nil {
			panic(err)
		}
		checkSame(data, result, "test")
	}))
	defer ts.Close()
	_, err := RESTRequest(ts.URL, &data, nil)
	if err != nil {
		panic(err)
	}
}

func TestRESTRequestResult(t *testing.T) {
	var data = map[string]string{
		"test": "value",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		reader := bytes.NewBuffer(jsonBytes)
		io.Copy(w, reader)
	}))
	defer ts.Close()
	var result = map[string]string{}
	_, err := RESTRequest(ts.URL, nil, &result)
	if err != nil {
		panic(err)
	}
	checkSame(data, result, "test")
}

func TestRESTRequestDataAndResult(t *testing.T) {
	var result = map[string]string{}
	var data = map[string]string{
		"test": "value",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, r.Body)
	}))
	defer ts.Close()
	_, err := RESTRequest(ts.URL, &data, &result)
	if err != nil {
		panic(err)
	}
	checkSame(data, result, "test")
}
