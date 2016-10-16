package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type DataObject interface {
	Save() DataObject
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello? Yes, this is dog!")
}

func ReturnJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func ReturnDataNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func createEntry(i interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewDecoder(r.Body).Decode(i); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(i); err != nil {
			panic(err)
		}
	}
}