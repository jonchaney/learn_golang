package main

import "fmt"

// Structs example
func Structs() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	v = Vertex{X: 2} // Y is initialized to zero
	fmt.Println(v)
	// struct pointers work how you would expect
	vertex := Vertex{3, 4}
	p := &vertex
	fmt.Println(p.X)
}
