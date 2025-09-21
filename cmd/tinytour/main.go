package main

import "fmt"

// main here holds a tiny tour of Go basics
// Highly recommend going through both
// - https://gobyexample.com/
// - https://go.dev/doc/effective_go
func main() {
	// Variables
	var hello string
	fmt.Println("var", hello) // "" which is the zeroed value of sting
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
	// - can't change their length unless you copy them to a bigger array
	b := [5]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3} // "..." lets the compiler figure out the length

	fmt.Printf("b: %d, c: %d\n", b, c) // b: [1 2 3 4 5], c: [1 2 3]

	// slices
	// - https://go.dev/blog/slices-intro and https://go.dev/doc/effective_go#slices
	// - slices point to an underlying array
	// - two slices can point to the same array
	var mySlice []int
	mySlice = append(mySlice, 1)
	fmt.Println(mySlice[0]) // 1

	mySlice2 := []int{1, 2, 3}
	fmt.Println(mySlice2[0]) // 1

	// makes a slice of length 5 with an underlying array capacity of 10
	// - the underlying array can be bigger than the slice because the slice is just a pointer to a view of the array
	// - if appending to a slice exceeds capacity, the slice is copied to another slice with double the capacity
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
// Capitalizing the first letter makes this public so other packages can use Plus
// NOTE! You can't export anything from package main so move Plus to a different package if you want to export it
func Plus(a, b int) int {
	return a + b
}
