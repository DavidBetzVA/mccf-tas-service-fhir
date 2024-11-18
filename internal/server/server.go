package server

import (
	"fmt"
	"net/http"

	"github.com/department-of-veterans-affairs/mccf-tas-go-shared/config"
	"github.com/gorilla/mux"
)

func CreateServer(addr string) *http.Server {
	router := setupRouter()
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/pnc", func(w http.ResponseWriter, r *http.Request) {
		searchByTagsHandler(w, r)
	}).Methods("GET")
	r.HandleFunc("/api/v1/pnc/{id}", func(w http.ResponseWriter, r *http.Request) {
		searchByIDHandler(w, r)
	}).Methods("GET")
	port := config.GetEnvIntDef("PORT", 8080)
	fmt.Printf("Server listening on port %d\n", port)
	return r
}
