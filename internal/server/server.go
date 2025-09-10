package server

import (
	"fmt"
	"gitignore/internal/gi"
	"gitignore/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start(conf utils.Config) {
	windowStart = time.Now()

	mux := mux.NewRouter()
	mux.Use(rateLimiter)

	mux.HandleFunc("/{catchall:.*}", gi.Generate).Methods("GET")

	fmt.Printf("Server running on port: %s", conf.PORT)
	if conf.ENV == "production" {
		http.ListenAndServeTLS(":"+conf.PORT, conf.CRT, conf.KEY, mux)
	} else {
		log.Fatal(http.ListenAndServe(":"+conf.PORT, mux))
	}
}
