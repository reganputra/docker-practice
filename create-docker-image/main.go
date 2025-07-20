package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/about", aboutHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page accessed")
	response := map[string]string{"message": "Welcome to the Home Page!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About page accessed")
	response := map[string]string{"message": "This is the About Page!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
