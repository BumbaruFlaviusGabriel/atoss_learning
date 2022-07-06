package main

import "fmt"

func main() {

	var name string
	var gender string
	var age int

	fmt.Scanln(&name)
	fmt.Scanln(&gender)
	fmt.Scanln(&age)

	var newPerson = person{
		name:   name,
		gender: gender,
		age:    age,
	}
	fmt.Print(newPerson)

}

type person struct {
	name   string
	gender string
	age    int
}
