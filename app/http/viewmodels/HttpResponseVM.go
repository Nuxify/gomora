package viewmodels

import (
	"encoding/json"
	"net/http"
)

// HTTPResponseVM base http viewmodel for json
type HTTPResponseVM struct {
	Code    int         `json:"code"`
	Records interface{} `json:"records"`
	Status  interface{} `json:"status"`
	Success bool        `json:"success"`
}

// JSON converts http responsewriter to JSON
func (response HTTPResponseVM) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}
