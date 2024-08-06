package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func returnResponseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		// Log Internal Server Errors explicitly to console
		log.Println("Responding with 5XX error:", message)
	}

	returnResponseAsJSON(w, code, ErrorResponse{
		Error: message,
	})
}

func returnResponseAsJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Internal Server Error: Json marshalling failed
		log.Printf("Failed to marshal payload to JSON: %v\n", err)
		fmt.Fprintf(w, "Failed to marshal payload to JSON: %v\n", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
