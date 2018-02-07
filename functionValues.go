package main

import "fmt"

// functions are values too and can be passed around just like other values

// FunctionValues example
func FunctionValues() {
	area := func(length, width float64) float64 {
		return length * width
	}

	square := Square{4, 5}

	fmt.Println(area(square.width, square.length))
	fmt.Println(calculateArea(square, area))
}

func calculateArea(s Square, cb func(length float64, width float64) float64) float64 {
	return cb(s.width, s.length)
}
