package services

import (
	"fmt"
	"learninggolanddb/db"
	"learninggolanddb/entity"
)

func GetAllCustomers() {
	customers := []entity.Customer{}
	db.GetDB().Where("name=?", "Carrefour").Find(&customers)
	fmt.Println(customers)
}

func DeleteCustomer(name string) {
	customers := entity.Customer{}
	db.GetDB().Where("name=?", name).Delete(&customers)
	fmt.Println(customers)
}

func CreateCustomer(customer entity.Customer) {
	db.GetDB().Create(&customer)
	fmt.Println(customer)
}
