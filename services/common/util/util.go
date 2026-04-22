package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return err
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(v)
	if err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	res := map[string]string{
		"error": err.Error(),
	}
	WriteJSON(w, statusCode, res)
}
