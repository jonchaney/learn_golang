package main

import (
	"fmt"
	"math"
)

// IfStatement example
func IfStatement() {
	fmt.Println("- if statement")
	// you can initialize a variable that is accessible in the if block only
	if value := math.Sqrt(25); value == 5 {
		fmt.Println("Square Root of 25 is 5!")
	} else {
		// value is available here too
		fmt.Printf("%v", value)
	}
}
