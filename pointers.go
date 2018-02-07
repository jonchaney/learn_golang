package main

import "fmt"

// Pointers example
func Pointers() {
	i := 96
	// declare a pointer to an integer
	var p *int
	// assign the address of i to pointer p
	p = &i
	// print the address
	fmt.Printf("%v\n", p)
	// dereference the pointer
	fmt.Printf("%v\n", *p)
}
