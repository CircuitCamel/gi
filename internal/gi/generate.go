package gi

import (
	"gitignore/internal/utils"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	chosen := mux.Vars(r)["catchall"]

	if chosen == "" {
		http.ServeFile(w, r, "templates/default.gitignore")
	} else if chosen == "list" {
		listing, err := utils.ListDir("templates")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(listing)
		return
	} else {
		w.Write(TemplateMaker(strings.Split(chosen, ",")))
		return
	}

	w.Write([]byte("not there"))
}
