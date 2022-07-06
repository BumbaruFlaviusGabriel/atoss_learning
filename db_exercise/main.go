package main

import (
	"db_exercise/db"
	"db_exercise/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var employee entity.Employee
var location entity.Location
var company entity.Company

func main() {
	db.InitDataBase()
	router := chi.NewRouter()
	//router.Get("/api/company", )
	router.Post("/api/company", CreateCompany)
	http.ListenAndServe(":8080", router)

}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &company)
}
