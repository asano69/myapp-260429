package main

import (
	"fmt"
	"myapp/internal/handler"
	"net/http"
	"os"
)

func runServer() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := host + ":" + port

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	fmt.Printf("listening on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
