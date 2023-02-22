package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "JSON")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "JSON")
	params := mux.Vars(r)
	book, _ := getBookById(params["id"])
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "JSON")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	_ = json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	book, _ := getBookById(params["id"])
	books = append(books, book)
	_ = json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for idx, book := range books {
		if book.Id == params["id"] {
			books = append(books[:idx], books[idx+1:]...)
		}
	}
	_ = json.NewEncoder(w).Encode(books)
}

func getBookById(id string) (Book, error) {
	for _, book := range books {
		if id == book.Id {
			return book, nil
		}
	}
	return Book{}, errors.New("not found")
}

func initBooks() {
	for i := 0; i < 10; i++ {
		book := Book{
			Id:       strconv.Itoa(i),
			Name:     strconv.Itoa(i),
			Category: "golang",
		}
		books = append(books, book)
	}
}

func main() {
	initBooks()
	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books/{id}", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("start server at 8080")
	log.Fatal(http.ListenAndServe("192.168.124.9:80", r))
}
