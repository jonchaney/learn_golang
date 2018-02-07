package main

import "fmt"

// HashTable example
func HashTable() {
	hash := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}

	fmt.Println(hash)
}
