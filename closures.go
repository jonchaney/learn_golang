package main

import "fmt"

// just like javascript, go functions can be closures
// adder function taken straight from golang tour
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Closures example
func Closures() {
	fn := adder()
	fmt.Println(fn(1))
	fmt.Println(adder()(2))
}
