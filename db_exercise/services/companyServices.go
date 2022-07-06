package services

import (
	"db_exercise/db"
	"db_exercise/entity"
	"fmt"
	"net/http"
)

func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	company := db.GetDB().Find(&entity.Company{})
	fmt.Println(w, company)
}
