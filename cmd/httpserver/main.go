package main

import (
	"goHttpServer/internal/controllers/healthcheck"
	"goHttpServer/internal/controllers/hello"
	"goHttpServer/pkg/utils"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("GET /health-check", healthcheck.Handle)
	server.HandleFunc("POST /hello", hello.Handle)

	err := utils.ListenAndServe(8080, server)
	log.Fatal(err)
}
