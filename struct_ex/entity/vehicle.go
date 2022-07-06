package entity

import "fmt"

type Vehicle interface {
	GetRide()
	GetMaxSpeed() int
	GetTotalTires() int
}

type Car struct {
	MaxSpeed   int
	TotalTires int
	Brand      string
}

type Bike struct {
	MaxSpeed   int
	TotalTires int
}

type Plane struct {
	MaxSpeed   int
	TotalTires int
}

//--------------------------------

func (c Car) GetRide() {
	fmt.Println("You can drive a car")
}

func (c Car) GetMaxSpeed() int {
	return c.MaxSpeed

}

func (c Car) GetTotalTires() int {
	return c.TotalTires
}

//--------------------------------

func (b Bike) GetRide() {
	fmt.Println("You can ride a bike")
}

func (b Bike) GetMaxSpeed() int {
	return b.MaxSpeed

}

func (b Bike) GetTotalTires() int {
	return b.TotalTires
}

//--------------------------------

func (p Plane) GetRide() {
	fmt.Println("You can fly a plane!")
}

func (p Plane) GetMaxSpeed() int {
	return p.MaxSpeed

}

func (p Plane) GetTotalTires() int {
	return p.MaxSpeed

}
