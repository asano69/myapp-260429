package main

import (
	"fmt"
	"myapp/internal/handlers"
	"net/http"
	"os"
)

type Mux map[string]http.HandlerFunc

func (m Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := m[r.URL.Path]; ok {
		h(w, r)
		return
	}
	http.NotFound(w, r)
}

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

	mux := Mux{
		"/":        handler.Index(),
		"/static/": handler.StaticHandler(),
	}

	fmt.Printf("listening on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
