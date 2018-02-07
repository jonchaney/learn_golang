package main

import "fmt"

// Maps example
func Maps() {
	// a map maps keys to values
	// the zero value of a map is nil
	// the make function returns a map of the given type, initialized and ready for use
	m := make(map[string]bool)
	m["Jonathan"] = true
	if m["Jonathan"] {
		fmt.Println("jonathan was here")
	}

	// map literal
	var literal = map[string]string{
		"Jon":   "Chaney",
		"Tamar": "Ariel",
	}

	var vertexLiteral = map[string]Vertex{
		"App Academy": {2, 3}, // you don't need to put the type here --> Vertex{2,3}
	}

	fmt.Println(literal)
	fmt.Println(vertexLiteral)
	// delete an element
	delete(literal, "Jon")
	fmt.Println(literal)
	// test if key is in m with two value assignment
	elem, ok := literal["Jon"]
	// if the key is not in the map, then elem is the zero value for the map's element type -- string's zero value is an empty string
	if !ok {
		fmt.Println("jon is not there!")
	} else {
		fmt.Printf("%v is there\n", elem)
	}
}
