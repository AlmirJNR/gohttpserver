package main

import (
	"goHttpServer/internal/controllers/healthcheck"
	"goHttpServer/internal/controllers/hello"
	"goHttpServer/internal/pkg"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("GET /health-check", healthcheck.Handle)
	server.HandleFunc("POST /hello", hello.Handle)

	err := pkg.ListenAndServe(8080, server)
	log.Fatal(err)
}
