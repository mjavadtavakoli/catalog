package server

import (
	"encoding/json"
	"net/http"
)

type Err struct {
	Message string
}

type Info struct {
	Message string
}

func Response(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Err{Message: err.Error()})
}

func InfoResponse(w http.ResponseWriter, message string) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Info{Message: message})
}

func Bind(w http.ResponseWriter, r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		ErrorResponse(w, err)
	}
	return err
}
