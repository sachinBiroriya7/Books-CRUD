package Routes

import (
	"Book-api-CRUD/Models"
	"net/http"

	"github.com/gorilla/mux"
)

func Deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for idx, item := range Models.Books {
		if item.Id == params["id"] {
			Models.Books = append(Models.Books[:idx], Models.Books[idx+1:]...)
			break
		}
	}
}
