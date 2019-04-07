package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s\n", r.URL.Path)
	fmt.Println("Server running")
	fmt.Println("Hello")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookTitle := vars["title"]
	pageNo := vars["page"]
	fmt.Fprintf(w, bookTitle+pageNo)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/books/{title}/page/{page}", getBook)

	http.ListenAndServe(":8080", r)
}
