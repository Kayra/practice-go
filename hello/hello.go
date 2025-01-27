package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Rie")
	fmt.Println(message)
}
