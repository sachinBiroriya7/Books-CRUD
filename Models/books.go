package Models

type Book struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Rating string  `json:"rating"`
	Author *Author `json:"Author"`
}

type Author struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

var Books []Book

func init() {
	// In Go, you can't place statements like append directly at the package level.
	//setting initial values to the Book slice inside the init function
	Books = append(Books, Book{Id: "1", Name: "Harry Potter", Rating: "4.7", Author: &Author{Fname: "J.K", Lname: "Rowling"}})
	Books = append(Books, Book{Id: "2", Name: "MahaBharat", Rating: "5", Author: &Author{Fname: "Ved", Lname: "Vyas"}})
}
