package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DavidBetzVA/mccf-tas-service-fhir/internal/server"
	"github.com/department-of-veterans-affairs/mccf-tas-go-shared/config"
)

func main() {
	port := config.GetEnvString("PORT", "3000")
	host := config.GetEnvString("HOST", "0.0.0.0")
	srv := server.CreateServer(fmt.Sprintf("%s:%s", host, port))
	fmt.Printf("Starting server at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
