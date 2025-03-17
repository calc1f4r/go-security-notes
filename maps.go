package main

func maps(){
	// Map declaration
	var a map[string]int
	b := make(map[string]int)
	c := map[string]int{"one": 1, "two": 2}
	
	// Accessing elements
	b["one"] = 1
	b["two"] = 2
	b["three"] = 3
	b["four"] = 4
	b["five"] = 5
	
	// Deleting elements
	delete(b, "three")
	
	// Checking if an element exists
	value, exists := b["three"]
	
	// Map length
	length := len(b)
	
	// Map iteration
	for key, value := range b {
		println(key, value)
	}
	
	// Map operations
	println("Map a:", a)
	println("Map b:", b)
	println("Map c:", c)
	println("Value of key 'one' in map b:", b["one"])
	println("Value of key 'three' in map b:", value, exists)
	println("Length of map b:", length)
}