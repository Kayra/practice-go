// Go is a pass-by-value language.
// An address is where a value is stored.
// To find an address of a variable, use the & operator before a variable.
// Pointers are a fundamental feature in Go that enable developers to manage memory efficiently and directly.
// Pointers are variables that store an address.
// A pointer is specific to what type of address it can store.
// The * operator can be used to assign a pointer the type of the value its address holds.
// The * operator can also be used to dereference a pointer and assign a new value.

package main

import "fmt"

func brainwash(saying *string) {
	*saying = "Beep Boop."
}

func addressesAndPointers() {
	greeting := "Hello there!"

	brainwash(&greeting)

	fmt.Println("greeting is now:", greeting)
}
