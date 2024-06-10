package utils

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func listen(port int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}

func ListenAndServe(port int, server *http.ServeMux) error {
	logger := log.Default()
	logger.Printf("Checking port %d availability...\n", port)
	listener, err := listen(port)
	for err != nil {
		logger.Printf("Port %d is busy\n", port)
		port += 1
		logger.Printf("Checking if port %d is also busy...\n", port)
		listener, err = listen(port)
	}

	err = listener.Close()
	if err != nil {
		logger.Panicf("Failed to close tcp listener on port %d\n", port)
	}
	logger.Printf("Starting server on port %d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), server)
}
