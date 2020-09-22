package main

import "fmt"

func main() {
	var numbers [10]int
	fmt.Println(numbers) // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(numbers)) // 10

	numbers[0] = 10

	newArray := [5]int{5, 10, 4, 5, 6}

	fmt.Println(newArray) // [5 10 4 5 6]
}