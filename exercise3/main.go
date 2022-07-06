package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Get("/welcome", handleWelcome)
	//http.HandleFunc("/welcome", handleWelcome) - we use router
	http.ListenAndServe(":8080", router)
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)

	name := r.URL.Query().Get("name")
	name = fmt.Sprintf("Welcome to our internship %s", name)
	fmt.Fprint(w, name)
}
