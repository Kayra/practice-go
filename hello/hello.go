package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings app: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Rie")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
