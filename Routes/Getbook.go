package Routes

import (
	"Book-api-CRUD/Models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, item := range Models.Books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
