// palindrome permutation: Given a string, write a function to check if it us a permutation
// of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards.
// A permutation is a rearrangement of letters. The palindrome does not need to be limited to
// just dictionary words

package main

import (
	"fmt"
	"strings"
)

func main() {
	if palindromePermutation("Tact Coa") {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

// create hash table and track number of odd appearancess
// O(n)
func palindromePermutation(str string) bool {
	odd := 0
	letters := make(map[string]int)
	for _, char := range str {
		if string(char) == " " {
			continue
		}
		letters[strings.ToLower(string(char))]++
		if letters[strings.ToLower(string(char))]%2 == 1 {
			odd++
		} else {
			odd--
		}
	}
	return odd <= 1
}
