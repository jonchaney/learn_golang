package main

import "fmt"

// a dynamically sized, flexible view into the elements on an array
// slices are a data structure built on top of an array
// "a slice is a descriptor of an array segement. It consists of a pointer to the array,
// the length of the segment (the slice), and its capacity (the underlying array)"

// Slices examples
func Slices() {
	// declare and initialze an array
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array", a)
	b := [...]int{1, 2} // have the compiler count the length for you with ...
	fmt.Println("array", b)
	// create a slice of the array a
	slice := a[1:4]
	fmt.Println("slice", slice)
	// a slice does not store any data, it just describes a section of the underlying array
	// changing the elements of a slice modifies the underlying array
	slice[0] = 9
	fmt.Println("modified array", a)

	// a slice has both a length and a capacity
	// length is the length of the slice
	fmt.Println("length of slice", len(slice))
	// capacity is the length of the underlying array
	fmt.Println("capcity of slice", cap(slice))

	// array literal
	literal := [3]bool{false, false, false}
	// slice literal -- this creates the same as above and then builds a slice that references it
	sliceLiteral := []bool{false, false, false}
	fmt.Println(literal, sliceLiteral)

	// the zero value of a slice is nil
	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("nil! slice")
	}

	// you can craete slices with the make function
	// this is how you make dynamically size arrays
	s := make([]int, 5)        // length=5 cap=5
	s2 := make([]string, 0, 5) // length=0 cap=5
	fmt.Println(s)
	fmt.Println(s2)

	// append to a slice
	s = []int{2, 3, 4}
	fmt.Println(s)
	s = append(s, 4)
	fmt.Println(s)

	// pointers to an array
	// unlike C, an array is NOT a pointer to the first
	// element in the array, an array variable denotes the ENTIRE array
	// you can make a pointer to the array, but that would be a pointer, not the array
	fmt.Println("POINTER TO ARRAY")
	array := []int{1, 2, 3}
	var ptr = &array
	fmt.Println(*ptr) // you cannot index into it

	// appending elements to a slice
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	y := x[:] // y now references the array x, it is not a copy, it is a new slice that points to the original array
	fmt.Println(x)
	fmt.Println(y)
	y[0] = "Bill"
	y = append(y, "Test") // append
	fmt.Println(x)
	fmt.Println(y)

	d := []string{"John", "Paul"}
	e := []string{"George", "Ringo", "Pete"}
	d = append(d, e...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println(d)
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}

	// a slice cannot be grown beyond its capacity, it will cause a runtime panic
	// to make a slice bigger, use the following code
	// append does this automatically
	w := []string{"Лайка", "Белка", "Стрелка"}
	t := make([]string, len(w), (cap(w)+1)*2)
	copy(t, w) // this copies all the elements in w to the new larger slice t
	w = t
	fmt.Println(cap(w))

}
