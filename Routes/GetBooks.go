package Routes

import (
	"Book-api-CRUD/Models"
	"encoding/json"
	"net/http"
)

func Getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Models.Books)
}
