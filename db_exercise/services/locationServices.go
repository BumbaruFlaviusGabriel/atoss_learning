package services

import (
	"db_exercise/db"
	"db_exercise/entity"
	"fmt"
	"net/http"
)

func GetAllLocations(w http.ResponseWriter, r *http.Request) {
	location := db.GetDB().Find(&entity.Location{})
	fmt.Println(w, location)
}
