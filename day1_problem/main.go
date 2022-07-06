package main

import (
	"day1_problem/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var employee entities.Employee
var location entities.Location
var company entities.Company

func main() {
	router := chi.NewRouter()
	router.Get("/api/welcome", handleWelcome)
	router.Get("/api/company", handleCompany)
	router.Post("/api/company", CreateCompany)
	http.ListenAndServe(":8080", router)
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)

	value := fmt.Sprintf("Welcome to Atoss")
	fmt.Fprint(w, value)
}

func handleCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, name)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &company)

	company = entities.Company{
		Employees: []entities.Employee{},
		Location:  "TM",
		Name:      "Atoss",
	}

}
