package main

import (
	"fmt"
	"learngo/workshop"
)

//1st letter small = private
//1st letter CAPS = public

func main() {
	fmt.Println("test")
	a, _ := add(1, 2) // _ ignores a var , := gets the type
	fmt.Println(a)
	workshop.Print("test\n")

	for test := 0; test < 10; test++ {
		fmt.Printf("I am number %v\n", test) //%v = value
	}

	people := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i := 0; i < len(people); i++ {
		fmt.Printf("I am person %v\n", people[i])
	}

	for i, _ := range people {
		fmt.Printf("I am person %v\n", people[i])
	}

	var newIntern = intern{
		name:   "Flavius",
		gender: "Male",
		age:    22,
	}

	fmt.Print(newIntern)
}

func add(a int, b int) (int, int) {
	return a + b, a - b
}

type intern struct {
	name   string
	gender string
	age    int
}
