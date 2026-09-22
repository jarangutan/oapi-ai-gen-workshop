package main

import "fmt"

// main here holds a tiny tour of Go basics
// Highly recommend going through both
// - https://gobyexample.com/
// - https://go.dev/doc/effective_go
func main() {
	// Variables
	var hello string
	fmt.Println("var", hello) // "" which is the zero value of string
	// var hello = "world"

	myNum := 13
	fmt.Println("walrus", myNum) // 13

	// https://go.dev/blog/constants
	const myconst string = "hello"
	fmt.Println("const", myconst) // "hello"

	// Loops
	for i := 0; i <= 2; i++ {
		fmt.Println("for loop", i) // for loop 0 -> for loop 2
	}

	for i := range 2 {
		fmt.Println("range loop", i) // range loop 0 -> range loop 1
	}

	// if/else if/else
	name1 := "bob"
	if name1 == "bob" {
		fmt.Println("is bob")
	} else if name1 == "jeff" {
		fmt.Println("is jeff")
	} else {
		fmt.Println("???")
	}

	// switch
	name2 := "bob"
	switch name2 {
	case "bob":
		fmt.Println("is bob")
	case "jeff":
		fmt.Println("is jeff")
	default:
		fmt.Println("???")
	}

	// arrays
	// - arrays have a fixed length
	// - copying the values into a larger array gives you a new array
	b := [5]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3} // "..." lets the compiler figure out the length

	fmt.Printf("b: %d, c: %d\n", b, c) // b: [1 2 3 4 5], c: [1 2 3]

	// slices
	// - https://go.dev/blog/slices-intro and https://go.dev/doc/effective_go#slices
	// - a slice describes part of an underlying array and tracks its length and capacity
	// - two slices can share the same underlying array
	var mySlice []int
	mySlice = append(mySlice, 1)
	fmt.Println(mySlice[0]) // 1

	mySlice2 := []int{1, 2, 3}
	fmt.Println(mySlice2[0]) // 1

	// makes a slice of length 5 with an underlying array capacity of 10
	// - the underlying array can have more room than the slice currently uses
	// - if append runs out of capacity, Go allocates another array and copies the values over
	// - always assign append's result back because it may use a new underlying array
	mySlice3 := make([]int, 5, 10)
	fmt.Println(mySlice3) // [0 0 0 0 0]
	mySlice3 = append(mySlice3, 1, 2, 3, 4, 5, 6)
	fmt.Println(mySlice3) // [0 0 0 0 0, 1, 2, 3, 4, 5, 6]

	// maps
	myMap := make(map[string]int) // or myMap := map[string]int{"k1": 1, "K2": 2}
	myMap["k1"] = 1
	myMap["k2"] = 2
	delete(myMap, "k1")
	fmt.Println("myMap", myMap) // myMap map[k2:2]

	// structs
	// are typed collections of fields useful for grouping data together into records
	type gopher struct {
		name   string
		isCute bool
	}

	myGopher := gopher{
		name:   "Gopherina",
		isCute: true,
	}
	fmt.Println("myGopher", myGopher) // myGopher {Gopherina true}

	// funcs
	answer := Plus(1, 2)
	fmt.Println("answer", answer) // 3
}

// Plus takes a and b and returns the sum
// Adding comments above things turns the comment into documentation. see https://tip.golang.org/doc/comment
//
// Capitalizing the first letter exports Plus
// NOTE! package main can't be imported, so move Plus to another package
// if you want to use it elsewhere
func Plus(a, b int) int {
	return a + b
}
