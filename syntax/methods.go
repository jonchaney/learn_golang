package main

import "fmt"

// MyFloat declare methods on types
type MyFloat float64

// Methods example
func Methods() {
	emp := Employee{"Jonathan Chaney", 1548898}
	emp.info()
	var a MyFloat
	a = -5
	fmt.Println(a.Abs())
}

// method attached to Employee type
func (e Employee) info() {
	fmt.Printf("name: %v\nID: %v\n", e.name, e.id)
}

// Abs exported method must have comment
// you cannot define methods with a reciever whose type
// is defined in another package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// methods with pointer recievers can modify the value to which the reciever points
