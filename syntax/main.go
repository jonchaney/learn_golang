package main

import (
	"fmt"
)

func main() {
	fmt.Println(functions("This function returns a string"))
	IfStatement()
	Loops()
	SwitchStatement()
	DeferStatement()
	Pointers()
	Structs()
	Arrays()
	Slices()
	HashTable()
	Ranges()
	Maps()
	FunctionValues()
	Closures()
	Methods()
}

// declare its return value and parameter type
func functions(str string) string {
	// declare a variable and initialize to zero
	var x float64
	var y float64
	x, y = twoReturnValues(1.5, 1.7)
	fmt.Printf("%v %v\n", x, y)
	return str
}

// for two return values state the type to be returned
func twoReturnValues(x, y float64) (float64, float64) {
	return x, y
}
