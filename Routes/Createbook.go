package Routes

import (
	"Book-api-CRUD/Models"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

func Createbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := rand.Intn(100000)

	var newBook Models.Book
	json.NewDecoder(r.Body).Decode(&newBook)
	newBook.Id = strconv.Itoa(id)

	Models.Books = append(Models.Books, newBook)

	json.NewEncoder(w).Encode(newBook)
}
