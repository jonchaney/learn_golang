package main

import "fmt"

// the defer statement defers the execution of a function until the surrounding function returns

// deferred function calls are pushed onto a stack. When a function returns,its deferred calls are
// executed in last-n-first-out order
// in what order will statements get printed in the following functions?

// DeferStatement example
func DeferStatement() {
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
