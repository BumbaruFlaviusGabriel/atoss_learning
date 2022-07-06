package main

import (
	"learninggolanddb/db"
	"learninggolanddb/entity"
	"learninggolanddb/services"
)

func main() {
	db.InitDataBase()
	services.GetAllCustomers()
	services.DeleteCustomer("Profi")
	customer := entity.Customer{
		Test:        "test",
		Address:     "test1",
		Description: "test2",
	}
	services.CreateCustomer(customer)
}
