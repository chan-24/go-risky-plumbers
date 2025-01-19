package risks

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

var (
	risks = make(map[string]Risk) // Map to store risks by UUID - internal DB for now
)

// Middleware function to ensure all requests and responses are in JSON format
func JsonContentTypeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request is JSON
		if r.Header.Get("Content-Type") != "application/json" && r.Method != http.MethodGet {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		// Ensure the response is JSON
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// Get all Risks
func GetRisks(w http.ResponseWriter, r *http.Request) {
	// Send empty list instead of null if no risks are present
	if len(risks) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode([]Risk{}); err != nil {
			http.Error(w, "Error Encoding", http.StatusInternalServerError)
			return
		} // Encode as an empty array
		return
	}

	var riskList []Risk
	for _, risk := range risks {
		riskList = append(riskList, risk)
	}

	// Return the risks as JSON
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(riskList); err != nil {
		http.Error(w, "Error Encoding", http.StatusInternalServerError)
		return
	}
}

// Create a new Risk
func CreateRisk(w http.ResponseWriter, r *http.Request) {
	var newRisk Risk

	// Decode the JSON body into the Risk struct
	if err := json.NewDecoder(r.Body).Decode(&newRisk); err != nil {
		http.Error(w, "Invalid Input data", http.StatusBadRequest)
		return
	}

	// check if mandatory field "state" is provided and valid
	if newRisk.State == "" {
		http.Error(w, "State can not be empty. Please provide valid state for the risk", http.StatusBadRequest)
		return
	}

	if !isValidState(newRisk.State) {
		http.Error(w, "Invalid State. Should be one of - open, closed, accepted, investigating", http.StatusBadRequest)
		return
	}

	// Generate new UUID for the risk
	newRisk.ID = uuid.New().String()

	// Store the new risk
	risks[newRisk.ID] = newRisk
	log.Printf("New Risk added %s", newRisk)

	// Return the created risk
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newRisk); err != nil {
		http.Error(w, "Error Encoding", http.StatusInternalServerError)
		return
	}
}

// Get Risk by ID
func GetRiskByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	id := r.URL.Path[len("/v1/risks/"):]

	// Check if the risk exists
	risk, exists := risks[id]
	if !exists {
		http.Error(w, "risk not found", http.StatusNotFound)
		return
	}

	// Return the risk as JSON
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(risk); err != nil {
		http.Error(w, "Error Encoding", http.StatusInternalServerError)
		return
	}

}
