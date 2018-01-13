package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	// a hash aka map
	m := make(map[string]bool)
	m["jon"] = true
	fmt.Println(m["jon"])

}
