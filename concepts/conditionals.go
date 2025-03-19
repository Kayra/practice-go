package main

import "fmt"

func conditionals() {

	number := 1738
	if number == 0 {
		fmt.Println("number is 0")
	} else if number%2 == 0 {
		fmt.Printf("number %d is even.\n", number)
	} else {
		fmt.Printf("number %d is odd.\n", number)
	}

	// go conditionals allow a small statement to be executed before the condition is evaluated
	i := 0
	if i++; i > 0 {
		fmt.Println("This conditional will be met.")
	}

}
