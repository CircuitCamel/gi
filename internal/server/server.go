package server

import (
	"gitignore/internal/gi"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()

	mux.HandleFunc("/{catchall:.*}", gi.Generate).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
