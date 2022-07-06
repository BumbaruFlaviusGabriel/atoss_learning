package db

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDataBase() {

	migrateConn, err := migrate.New("file://db/migrate", "postgres://postgres:mysecretpassword@localhost:5432/learning?sslmode=disable")

	if err != nil {
		fmt.Println(err)
		return
	}

	version, _, _ := migrateConn.Version()
	fmt.Println(version)
	if version != 1 {
		migrateConn.Migrate(1)
	}
	migrateConn.Close()
	dbConn, err := gorm.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/learning?sslmode=disable")

	db = dbConn
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}
