package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv" 
	"github.com/gorilla/mux"
)


// Book struct (Model)
type Book struct {
	ID 		string 	`json:"id"`
	Isbn 	string 	`json:"isbn"`
	Title 	string 	`json:"title"`
	Author 	*Author  `json:"author"`
	
}


// GET Al Books
func getBooks(w http.ResponseWriter, r *http.Request) {


}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	

}

//Create Book
func createBook(w http.ResponseWriter, r *http.Request) {
	

}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	

}


// Author Struct
type Author struct {
	Firstname string `json:"firstname "`
	Lastname   string `json:"lastname  "`
}

func main() {
	// Init Router

	r := mux.NewRouter()

	// route handlers /endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")


	//server
	log.Fatal(http.ListenAndServe(":8000", r))


	
}