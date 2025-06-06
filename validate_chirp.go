package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Body string `json:"body"`
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	type validResponse struct {
		Valid bool `json:"valid"`
	}
	// Validate the request method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp, _ := json.Marshal(errorResponse{Error: "Method not allowed"})
		w.Write(resp)
		return
	}
	decoder := json.NewDecoder(r.Body)
	body := requestBody{}
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(errorResponse{Error: "Invalid JSON"})
		w.Write(resp)
		return
	}
	// Chirp length validation (assuming 140 chars max)
	if len(body.Body) > 140 {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(errorResponse{Error: "Chirp is too long"})
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(validResponse{Valid: true})
	w.Write(resp)
}
