package main

import "fmt"

// Arrays demonstrates various array operations in Go
func Arrays() {
	// Fixed-size array
	var a [5]int
	fmt.Println("Empty array:", a)
	
	// Initialize with values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", b)
	
	// Array with inferred size
	c := [...]int{1, 2, 3}
	fmt.Println("Array with inferred size:", c)
	
	// Accessing elements
	fmt.Println("First element:", b[0])
	
	// Modifying elements
	b[0] = 10
	fmt.Println("After modification:", b)
	
	// Array length
	fmt.Println("Array length:", len(b))
	

}



// different between array and slice
// array is a fixed size, slice is a dynamic size
// array is a value type, slice is a reference type

func arrayAndSlice() {
	// array	
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	// slice
	slice := a[0:2] 

	fmt.Println(a)
	fmt.Println(slice)
	slice = append(slice, 6)
	fmt.Println(slice)
}
