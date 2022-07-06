package main

import (
	"fmt"
	"golang_pointers/entity"
)

func main() {
	a := "test"
	c := 2
	d := &a
	f := &c
	fmt.Println(*d, *f) // sau &a, &c
	passbyValue(c)
	fmt.Println("Value of C after passbyValue", c)
	passbyReference(&c)
	fmt.Println("Value of C after passbyReference", c)
	dog := entity.Dog{
		Name:  "Tom",
		Breed: "German Shephard",
	}
	canyouWalk(dog)

}

func passbyValue(c int) {
	c = 4
}

func passbyReference(c *int) {
	*c = 5
}

func canyouWalk(animal entity.Animal) {
	dog, ok := animal.(entity.Dog)
	fmt.Print(ok)
	if ok {
		fmt.Println("I am called a dog")
		dog.Bite()
	}
	dog.Bite()
	animal.Walk()
}
