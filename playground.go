package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(functions("This function returns a string"))
	ifStatement()
	loops()
	switchStatement()
	deferStatement()
	pointers()
	structs()
	arrays()
	slices()
	hashTable()
	ranges()
}

// declare its return value and parameter type
func functions(str string) string {
	// declare a variable and initialize to zero
	var x float64
	var y float64
	x, y = twoReturnValues(1.5, 1.7)
	fmt.Printf("%v %v\n", x, y)
	return str
}

// for two return values state the type to be returned
func twoReturnValues(x, y float64) (float64, float64) {
	return x, y
}

func ifStatement() {
	fmt.Println("- if statement")
	// you can initialize a variable that is accessible in the if block only
	if value := math.Sqrt(25); value == 5 {
		fmt.Println("Square Root of 25 is 5!")
	} else {
		// value is available here too
		fmt.Printf("%v", value)
	}
}

func loops() {
	// while is spelled for in golang
	fmt.Println("- while loop")
	i := 0
	for i < 10 {
		fmt.Printf("%v", i)
		i++
	}
	// for loop
	fmt.Println("\n- for loop")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", i)
	}
}

func switchStatement() {
	fmt.Println("\n- switch statement")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// without a condition, the switch acts like a clean way to write long if-else statements
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

// the defer statement defers the execution of a function until the surrounding function returns
// deferred function calls are pushed onto a stack. When a function returns,its deferred calls are
// executed in last-n-first-out order
// in what order will statements get printed in the following functions?
func deferStatement() {
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

func pointers() {
	i := 96
	// declare a pointer to an integer
	var p *int
	// assign the address of i to pointer p
	p = &i
	// print the address
	fmt.Printf("%v\n", p)
	// dereference the pointer
	fmt.Printf("%v\n", *p)

}

// Vertex - a struct is collection of fields
type Vertex struct {
	X int
	Y int
}

func structs() {
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

func arrays() {
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

// a dynamically sized, flexible view into the elements on an array
// slices are a data structure built on top of an array
// "a slice is a descriptor of an array segement. It consists of a pointer to the array,
// the length of the segment (the slice), and its capacity (the underlying array)"
func slices() {
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

func hashTable() {
	hash := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}

	fmt.Println(hash)
}

func ranges() {
	name := []string{"jon", "chaney"}
	// if an index is not used, put an underscore in its place
	for _, n := range name {
		fmt.Println(n)
	}
	// if you only want the index, drop the value entirely!
	for i := range name {
		fmt.Println(name[i])
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println(pow)
}
