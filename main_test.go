package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Create request for handler to be tested with.
	req, _ := http.NewRequest("GET", "/check", nil)

	// Response recorder, to record response.
	rr := httptest.NewRecorder()

	// Attach handler to function and serve to handle routes.
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("Wrong status code, success not acheived, health unknown!")
	}

	expected := "{\"alive\":true}"
	if expected != rr.Body.String() {
		t.Errorf("The result %v is incorrect, we got %v back instead!", expected, rr.Body.String())
	}
}

func TestErrorResponseHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/errexample", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ErrorResponseHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Error("This was supposed to be a 500 error! WTF?!?!?")
	}
}

func TestResponseExampleHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/response", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ResponseExampleHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("Well that just smarts!")
	}

	if !strings.HasPrefix(rr.Body.String(), "Testing") {
		t.Errorf("The string was not prefixed as expected. Got %v.", rr.Body.String())
	}
}
