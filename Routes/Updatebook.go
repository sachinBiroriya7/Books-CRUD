package Routes

import (
	"Book-api-CRUD/Models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Updatebook(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type before writing any response
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	bookID := param["id"]

	var updateBook Models.Book
	// Decode the request body into updateBook
	if err := json.NewDecoder(r.Body).Decode(&updateBook); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Find and update the book
	for i, item := range Models.Books {
		if item.Id == bookID {
			Models.Books[i].Name = updateBook.Name
			Models.Books[i].Rating = updateBook.Rating
			Models.Books[i].Author.Fname = updateBook.Author.Fname
			Models.Books[i].Author.Lname = updateBook.Author.Lname

			// Return the updated book as a response
			json.NewEncoder(w).Encode(Models.Books[i])
			return
		}
	}

	// If book is not found, return a 404 error
	http.Error(w, "Book not found", http.StatusNotFound)
}
