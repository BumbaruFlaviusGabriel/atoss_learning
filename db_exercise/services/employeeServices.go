package services

import (
	"db_exercise/db"
	"db_exercise/entity"
	"fmt"
	"net/http"
)

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	employee := db.GetDB().Find(&entity.Employee{})
	fmt.Println(w, employee)
}
