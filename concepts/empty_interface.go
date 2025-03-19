// An empty interface can hold values of any type
// This can be helpful when working with values that can have multiple types
// Empty interfaces can be risky as they override compile-time type checking
// Using type assertions and type switches can help manage empty interfaces

package main

import "fmt"

type emptyInter interface{}

func processData(data emptyInter) {
	switch v := data.(type) {
	case string:
		fmt.Println("Processing string:", v)
	case int:
		fmt.Println("Processing int:", v*2)
	case bool:
		fmt.Println("Processing bool:", !v)
	default:
		fmt.Println("Unknown type, cannot process")
	}
}

// The any syntax can also be used to achieve the same result
func logData(data any) {
	fmt.Println("Logging data:", data)
}

func emptyInterface() {

	processData("Test")
	processData(7357)
	processData(true)
	processData(2.4)

	logData("Test")
	logData(7357)

}
