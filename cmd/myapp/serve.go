package main

import (
	"fmt"
	"myapp/internal/db"
	"myapp/internal/handler"
	"net/http"
	"os"
	"path/filepath"
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

	dataDir := os.Getenv("DATA")
	if dataDir == "" {
		dataDir = "data"
	}

	database, err := db.Open(filepath.Join(dataDir, "sqlite3.db"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "db error: %v\n", err)
		os.Exit(1)
	}
	defer database.Close()

	addr := host + ":" + port
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, database, dataDir)

	fmt.Printf("listening on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
