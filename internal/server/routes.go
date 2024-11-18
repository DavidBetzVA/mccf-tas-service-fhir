package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func searchByTagsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("searchByTagsHandler")
	w.Header().Set("Content-Type", "application/json")
	var obj interface{}
	jsonBuffer, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, string(jsonBuffer))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func searchByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("searchByIDHandler")
	w.Header().Set("Content-Type", "application/json")
	var obj interface{}
	jsonBuffer, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, string(jsonBuffer))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
