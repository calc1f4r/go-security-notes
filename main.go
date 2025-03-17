package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Go Basics Demonstrations ===")
	
	// Run basic demonstrations
	fmt.Println("\n--- Basic Demonstrations ---")
	variableShows()
	
	fmt.Println("\n--- Function Demonstrations ---")
	fmt.Println("add(1, 2) =", add(1, 2))
	fmt.Println("meow() =", meow())
	
	fmt.Println("\n--- Error Handling ---")
	// Error handling example
	result, err := division(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	
	// Try with valid inputs
	if result, err := division(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	
	fmt.Println("\n--- Array Demonstrations ---")
	Arrays()
	
	fmt.Println("\n--- Array and Slice Demonstrations ---")
	arrayAndSlice()
	
	fmt.Println("\n--- Map Demonstrations ---")
	maps()
} 