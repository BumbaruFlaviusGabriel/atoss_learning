package entity

import "fmt"

type Animal interface {
	Walk()
	Talk(s string) string
}

type WildAnimal interface {
	Bite()
}

//--------------------------------

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Walk() {
	fmt.Println("Yes, I can walk")
}

func (d Dog) Talk(s string) string {
	return fmt.Sprintf("%s baw baw baw", s)
}

func (d Dog) Bite() {
	fmt.Println("I can bite!")
}

//--------------------------------

type Cat struct {
	Name string
}

func (c Cat) Walk() {
	fmt.Println("No, I'm too lazy to walk")
}

func (c Cat) Talk(s string) string {
	return fmt.Sprintf("%s %s", s, c.Name)
}
