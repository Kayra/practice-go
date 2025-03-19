// Programming to an interface rather than specific implementations is done using generic interfaces instead of concrete types.

// Some benefits of programming to an interface:
// Adaptability: Keeps the code less dependent on specific implementations.
// Code Reuse: Reduces complexity and duplication.
// Expansion and Flexibility: Enables easier updates and growth without major overhauls.

package main

import "fmt"

// Interface defining methods required for vehicles
type Vehicle interface {
	CalculateRentalCost() int
}

// Scooter struct for storing scooter usage data
type Scooter struct {
	hours int
}

func (s Scooter) CalculateRentalCost() int {
	const hourlyRate int = 5
	return s.hours * hourlyRate
}

// Car struct for storing car usage data
type Car struct {
	hours int
	miles int
}

func (c Car) CalculateRentalCost() int {
	const hourlyRate int = 10
	const mileageRate int = 5
	return (c.hours * hourlyRate) + (c.miles * mileageRate)
}

// This function is an implicit method that will print the rental cost for any vehicle that implements the Vehicle interface.
func PrintRentalCost(v Vehicle) {
	fmt.Printf("Rental Cost: $%d\n", v.CalculateRentalCost())
}

func interfaces() {
	var vehicleOne Vehicle = Scooter{hours: 4}
	PrintRentalCost(vehicleOne)

	var vehicleTwo Vehicle
	vehicleTwo = Car{hours: 2, miles: 4}
	PrintRentalCost(vehicleTwo)
}
