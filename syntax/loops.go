package main

import "fmt"

// Loops example
func Loops() {
	// while is spelled for in golang
	fmt.Println("- while loop")
	i := 0
	for i < 10 {
		fmt.Printf("%v", i)
		i++
	}
	// for loop
	fmt.Println("\n- for loop")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", i)
	}
}
