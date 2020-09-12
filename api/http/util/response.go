package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	schemaSuccess struct {
		Message string `json:"message"`
	}
	schemaError struct {
		Error string `json:"error"`
	}
)

func Response(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	sch := schemaSuccess{
		Message: msg,
	}
	js, err := json.Marshal(&sch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		log.Printf("failed to write a resonse: %v", err)
	}
}

func Error(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	sch := schemaError{
		Error: msg,
	}
	js, err := json.Marshal(&sch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		log.Printf("failed to write a resonse: %v", err)
	}
}
