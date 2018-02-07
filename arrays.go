package main

import "fmt"

func Arrays() {
	// similar to C++, a size must be declared
	var a [3]string
	for i := 0; i < len(a); i++ {
		a[i] = "Jon"
		fmt.Printf("%v", a[i])
	}

	// declare and initialize values
	b := [3]int{1, 2, 3}

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

}
