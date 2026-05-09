package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	req:=httptest.NewRequest(
		http.MethodGet,
		"/hello",
		nil,
	)

	rec:=httptest.NewRecorder()

	handler:=http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rec, req)

	if rec.Code !=http.StatusOK{
		t.Errorf("Expected status 200 but got %d ", rec.Code)

	}

	expected:="hello"

	if rec.Body.String() != expected{
		t.Errorf("expected %s but got %s", expected, rec.Body.String())
	}

}