package gi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	chosen := mux.Vars(r)["catchall"]

	if chosen == "" {
		http.ServeFile(w, r, "templates/default.gitignore")
	}

	w.Write([]byte("not there"))
}
