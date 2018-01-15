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

// create a hash counter - letters as keys and number it appears as values
// if more than one letter appears an odd number of times, no palindrome exists
// O(n)
func palindromePermutation(str string) bool {
	// create hash table
	// O(n)
	letters := make(map[string]int)
	for _, r := range str {
		if string(r) == " " {
			continue
		}
		letters[strings.ToLower(string(r))]++
	}
	// check if there is more than one letter that appears an odd
	// number of times
	// O(m) -- length of hash counter
	odd := 0
	for _, r := range str {
		if letters[strings.ToLower(string(r))]%2 != 0 {
			odd++
		}
		if odd > 1 {
			return false
		}
	}
	return true
}
