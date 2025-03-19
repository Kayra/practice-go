package main

import "fmt"

type Maintainer interface {
	Service(vehicleID int) string
}

// In Go, methods are bound to receiver types.
// Receivers can be value or pointer types, influencing how methods interact with the data.

// Receivers:
// appear in their own argument list between the func keyword and the method name, similar to the this or self keyword in other languages
// are typically named with one or two letters
// determine interface implementations separately based on their type: value receivers or pointer receivers

type MaintenanceService struct {
	location string
}

type CleaningService struct {
	waterGallons int
}

// Value receivers operate on copies of the data. They do not modify the original instance.
func (ms MaintenanceService) Service(vehicleID int) string {
	location := ms.location
	return fmt.Sprintf("Vehicle %d has been serviced at %s.", vehicleID, location)
}

// Pointer receivers can modify the original data instance because they reference the memory address.
func (cs *CleaningService) Service(vehicleID int) string {
	cs.waterGallons--
	return fmt.Sprintf("Vehicle %d has been serviced with %d gallons of water remaining.", vehicleID, cs.waterGallons)
}

func receivers() {

	// Go does not require an explicit declaration to implement an interface.
	// Go considers it implemented as long as a type has all the methods specified in the interface.
	ms := MaintenanceService{location: "Downtown Garage"}
	cs := &CleaningService{waterGallons: 10}

	var maintainer Maintainer

	maintainer = ms
	fmt.Println(maintainer.Service(1))

	maintainer = cs
	fmt.Println(maintainer.Service(2))
	fmt.Printf("Global scope water gallons remaining is %d.\n", cs.waterGallons)

}
