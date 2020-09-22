package main

import "fmt"

func main() {
	
	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

	fmt.Println(m)
	
	delete(m, "c")

	fmt.Println(m["c"])

	_, exists := m["a"]
	fmt.Println(exists)
	
	value, exists := m["b"]
	fmt.Println(value)

	// var x := map[string]int{}

	// newMap := map[string]int{}
}