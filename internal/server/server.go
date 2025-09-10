package server

import (
	"fmt"
	"gitignore/internal/gi"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	windowStart = time.Now()

	mux := mux.NewRouter()
	mux.Use(rateLimiter)

	mux.HandleFunc("/{catchall:.*}", gi.Generate).Methods("GET")

	fmt.Printf("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
