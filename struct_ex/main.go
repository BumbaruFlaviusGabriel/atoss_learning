package main

import "struct_ex/entity"

func main() {

	car := entity.Car{
		MaxSpeed:   240,
		TotalTires: 4,
		Brand:      "Lambo",
	}

	bike := entity.Bike{
		MaxSpeed:   20,
		TotalTires: 2,
	}

	plane := entity.Plane{
		MaxSpeed:   700,
		TotalTires: 6,
	}
	canyouRide(car)
	canyouRide(bike)
	canyouRide(plane)

	maxSpeed(car)
	maxSpeed(bike)
	maxSpeed(plane)

	totalTires(car)
	totalTires(bike)
	totalTires(plane)

}

func canyouRide(vehicle entity.Vehicle) {
	vehicle.GetRide()
}

func maxSpeed(vehicle entity.Vehicle) {
	vehicle.GetMaxSpeed()
}

func totalTires(vehicle entity.Vehicle) {
	vehicle.GetTotalTires()
}
