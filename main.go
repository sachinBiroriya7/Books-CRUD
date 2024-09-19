package main

import (
	"Book-api-CRUD/Routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/books", Routes.Getbooks).Methods("GET")
	r.HandleFunc("/books/{id}", Routes.Getbook).Methods("GET")
	r.HandleFunc("/books", Routes.Createbook).Methods("POST")
	r.HandleFunc("/books/{id}", Routes.Updatebook).Methods("PUT")
	r.HandleFunc("/books/{id}", Routes.Deletebook).Methods("DELETE")

	fmt.Println("Starting server at PORT:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error in starting server :%v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Books Home Page")
}
