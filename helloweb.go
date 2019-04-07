package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// User is a representation of a user
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func decode(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "%s %s is %d years old!", user.FirstName, user.LastName, user.Age)
}

func encode(w http.ResponseWriter, r *http.Request) {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}
	json.NewEncoder(w).Encode(&user)
}

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
	r.HandleFunc("/encode", encode)
	r.HandleFunc("/decode", decode)
	r.HandleFunc("/books/{title}/page/{page}", getBook)

	http.ListenAndServe(":8080", r)
}
