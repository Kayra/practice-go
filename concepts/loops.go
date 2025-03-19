package main

import "fmt"

func loops() {

	fmt.Print("For looping through:")
	for number := 1; number < 6; number++ {
		fmt.Print(" ", number)
	}

	fmt.Println()

	fmt.Print("While looping through:")
	number := 1
	for number < 6 {
		fmt.Print(" ", number)
		number++
	}

	fmt.Println()

	fmt.Print("'Infinate' looping through:")
	infNumber := 1
	for {
		fmt.Print(" ", infNumber)
		infNumber++
		if infNumber > 5 {
			break
		}
	}

	fmt.Println()

	animals := []string{"Cat", "Dog", "Fish", "Turtle"}
	fmt.Printf("Finding dog in %v: ", animals)
	for index := 0; index < len(animals); index++ {
		if animals[index] == "Dog" {
			fmt.Printf("Found %s!", animals[index])
		}
	}

	fmt.Println()

	fmt.Println("Eating jelly beans: ")
	jellybeans := []string{"green", "blue", "yellow", "red", "green", "yellow", "rotten", "red"}
	for index := 0; index < len(jellybeans); index++ {
		if jellybeans[index] == "green" {
			continue
		}
		if jellybeans[index] == "rotten" {
			fmt.Print("A ", jellybeans[index], " jellybean, maybe that's enough for now...")
			break
		}
		fmt.Println("Ate the", jellybeans[index], "jellybean")
	}

	fmt.Println()

	fmt.Print("Range looping through array:")
	letters := []string{"A", "B", "C", "D", "E", "F", "G"}
	for _, value := range letters {
		fmt.Print(" ", value)
	}

	fmt.Println()

	fmt.Println("Range looping through map:")
	addressBook := map[string]string{
		"John":   "12 Main St",
		"Janet":  "56 Pleasant St",
		"Jordan": "88 Liberty Ln",
	}
	for key, value := range addressBook {
		fmt.Println("Name:", key, "Address:", value)
	}

}
