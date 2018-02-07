package main

import "fmt"

// Ranges example
func Ranges() {
	name := []string{"jon", "chaney"}
	// if an index is not used, put an underscore in its place
	for _, n := range name {
		fmt.Println(n)
	}
	// if you only want the index, drop the value entirely!
	for i := range name {
		fmt.Println(name[i])
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println(pow)
}
