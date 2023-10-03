package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/about", AboutHandler)

	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page!")
}
