package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(functions("This function returns a string"))
	ifStatement()
	loops()
	switchStatement()
	deferStatement()
	pointers()
	structs()
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

func ifStatement() {
	fmt.Println("- if statement")
	// you can initialize a variable that is accessible in the if block only
	if value := math.Sqrt(25); value == 5 {
		fmt.Println("Square Root of 25 is 5!")
	} else {
		// value is available here too
		fmt.Printf("%v", value)
	}
}

func loops() {
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

func switchStatement() {
	fmt.Println("\n- switch statement")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// without a condition, the switch acts like a clean way to write long if-else statements
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

// the defer statement defers the execution of a function until the surrounding function returns
// deferred function calls are pushed onto a stack. When a function returns,its deferred calls are
// executed in last-n-first-out order
// in what order will statements get printed in the following functions?
func deferStatement() {
	defer fmt.Println("World!")
	fmt.Printf("%v", "Hello, ")
	deferStack()
}

// defer statements are pushed onto a stack
// once the function is done executing they are popped off the stack
func deferStack() {
	fmt.Println("\nstart function")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v\n", i)
	}
	fmt.Println("end function")
}

func pointers() {
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

// Vertex - a struct is collection of fields
type Vertex struct {
	X int
	Y int
}

func structs() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
