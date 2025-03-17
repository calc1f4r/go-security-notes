package main

import (
	"errors"
	"fmt"
)

func variableShows() {
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println(c)

	// multiple variables
	var d, e, f = 1, 2, 3
	fmt.Println(d, e, f)

	// multiple variables with different types
	var (
		g int     = 1
		h string  = "hello"
		i float64 = 1.2
	)
	fmt.Println(g, h, i)

	// short declaration
	j := 10
	k := "hello"
	l := 1.2
	fmt.Println(j, k, l)

	// multiple short declaration
	m, n, o := 1, "hello", 1.2
	fmt.Println(m, n, o)

	// constant 
	const pi = 3.14
	fmt.Println(pi)

	// constant with different types
	const x, y, z = 1, "hello", 1.2
	fmt.Println(x, y, z)
}

func add(x int, y int) int {
	return x + y
}

func meow() string {
	return "meow"
}

// division returns the result of x divided by y
// returns error if y is zero
func division(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
} 