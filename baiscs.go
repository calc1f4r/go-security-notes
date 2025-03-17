package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(meow())
	fmt.Println(division(10, 0))
}

func add(x int, y int) int {
	return x + y
}

func meow() string {
	return "meow"
}


// check if it panics 
func division(x int, y int) int {
	return x / y
}
