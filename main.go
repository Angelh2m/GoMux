package main

import (
	"encoding/json"
	_"encoding/json"
	"github.com/gorilla/mux"
	_"log"
	"net/http"
	_"net/http"
	_"math/rand"
	_"strconv"
)

//Book struct (MODEL)

type Book struct {
	ID  string `json:"id"`
	Isbn  string `json:"isbn"`
	Title  string `json:"title"`
	Author  *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books []Book


func main()  {

	books = append(books,
		Book{ID: "1", Isbn: "213213", Title: "Book One", Author: &Author{ Firstname: "John", Lastname: "Doe",}})

	books = append(books,
			Book{ID: "2", Isbn: "4444444", Title: "Book Two", Author: &Author{ Firstname: "John", Lastname: "Doe",}})

	r := mux.NewRouter()
	//r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	//r.HandleFunc("/articles", ArticlesHandler)

	http.ListenAndServe(":8080", r)
	
	http.Handle("/", r)


}



func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {

}

func createBook(writer http.ResponseWriter, request *http.Request) {

}

func updateBook(writer http.ResponseWriter, request *http.Request) {

}

func deleteBook(writer http.ResponseWriter, request *http.Request) {

}
