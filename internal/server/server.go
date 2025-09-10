package server

import (
	"gitignore/internal/gi"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	windowStart = time.Now()

	mux := mux.NewRouter()
	mux.Use(rateLimiter)

	mux.HandleFunc("/{catchall:.*}", gi.Generate).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
