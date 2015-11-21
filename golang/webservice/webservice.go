package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/william", William)
	router.HandleFunc("/dad", Dad)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please select a target")
}

func William(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You chose William. Bad choice. Bad.")
}

func Dad(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You chose Dad. Excellent choice.")
}
