package risks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmptyGetRisks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/risks", nil)
	rr := httptest.NewRecorder()

	GetRisks(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", rr.Code)
	}

	expected := `[]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("expected %q, got %q", expected, rr.Body.String())
	}
}

// Post a risk and Get same risk by UUID
func TestPostRisks(t *testing.T) {

	jsonData := `{
		"state": "open",
		"title": "r1",
		"description": "low risk"
	}`

	// Create a new HTTP request with the JSON payload
	req := httptest.NewRequest(http.MethodPost, "/v1/risks", bytes.NewReader([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRisk)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status %v, got %v", http.StatusCreated, rr.Code)
	}

	var createdRisk Risk
	err := json.NewDecoder(rr.Body).Decode(&createdRisk)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	riskID := createdRisk.ID

	// test Get risk by UUID
	getReq := httptest.NewRequest(http.MethodGet, "/v1/risks/"+riskID, nil)
	getRR := httptest.NewRecorder()

	handlerGet := http.HandlerFunc(GetRiskByID)
	handlerGet.ServeHTTP(getRR, getReq)

	if getRR.Code != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, getRR.Code)
	}

	var fetchedRisk Risk
	err = json.NewDecoder(getRR.Body).Decode(&fetchedRisk)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	if fetchedRisk.ID != riskID {
		t.Errorf("expected ID %v, got %v", riskID, fetchedRisk.ID)
	}
	if fetchedRisk.Title != "r1" {
		t.Errorf("expected Name %v, got %v", "r1", fetchedRisk.Title)
	}
}

func TestInvalidPostRisks(t *testing.T) {

	jsonData := `{
		"state": "invalid",
		"title": "r1",
		"description": "low risk"
	}`

	// Create a new HTTP request with the JSON payload
	req := httptest.NewRequest(http.MethodPost, "/v1/risks", bytes.NewReader([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRisk)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
	}

	jsonData = `{
		"title": "r1",
		"description": "low risk"
	}`

	// Create a new HTTP request with the JSON payload
	req = httptest.NewRequest(http.MethodPost, "/v1/risks", bytes.NewReader([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(CreateRisk)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
	}
}
