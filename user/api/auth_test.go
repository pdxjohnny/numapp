// Make sure that only login and register are exposed
package api

import (
	"os"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/pdxjohnny/numapp/variables"
)

func TestMain(m *testing.M) {
	os.Setenv(variables.EnvTokenSignKey, "../../"+variables.TokenSignKeyDefault)
	os.Setenv(variables.EnvTokenVerifyKey, "../../"+variables.TokenVerifyKeyDefault)
	err := variables.LoadTokenKeys()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestLoginNeedsNoAuth(t *testing.T) {
	handler := MakeHandler()

	req := test.MakeSimpleRequest(
		"POST",
		"http://localhost"+variables.APIPathLoginUserServer,
		nil,
	)
	req.Header.Set("Accept-Encoding", "application/json")
	res := test.RunRequest(t, *handler, req)
	// Should error beacuse we didnt send a json
	res.BodyIs("{\"Error\":\"JSON payload is empty\"}")
	res.CodeIs(401)
	res.ContentTypeIsJson()
}

func TestRegisterNeedsNoAuth(t *testing.T) {
	handler := MakeHandler()

	req := test.MakeSimpleRequest(
		"POST",
		"http://localhost"+variables.APIPathRegisterUserServer,
		nil,
	)
	req.Header.Set("Accept-Encoding", "application/json")
	res := test.RunRequest(t, *handler, req)
	// Should error beacuse we didnt send a json
	res.BodyIs("{\"Error\":\"JSON payload is empty\"}")
	res.CodeIs(500)
	res.ContentTypeIsJson()
}

func TestBackendTokenHasAccess(t *testing.T) {
	handler := MakeHandler()

	req := test.MakeSimpleRequest(
		"POST",
		"http://localhost/",
		nil,
	)
	req.Header.Set("Accept-Encoding", "application/json")
	req.Header.Set("Authorization", "Bearer "+variables.BackendToken)
	res := test.RunRequest(t, *handler, req)
	// Should error beacuse we didnt send a json
	res.BodyIs("{\"Error\":\"Resource not found\"}")
	res.CodeIs(404)
	res.ContentTypeIsJson()
}

func TestAllOtherPathsNeedAuth(t *testing.T) {
	handler := MakeHandler()

	req := test.MakeSimpleRequest("GET", "http://localhost/", nil)
	res := test.RunRequest(t, *handler, req)
	res.CodeIs(401)
	res.ContentTypeIsJson()
}
