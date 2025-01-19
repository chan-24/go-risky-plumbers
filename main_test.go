package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	rootHandler(rr, req)

	// Check status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", rr.Code)
	}

	// Check response body
	expected := "Welcome to the Risky Plumbers Home Page! \nUse /v1/risks for accessing risks or creating new risks!"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}
