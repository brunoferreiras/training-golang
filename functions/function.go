package main

import "fmt"

func plus(a int, b int) int {

    return a + b
}

func plusPlus(a, b, c int) int {
    defer fmt.Println("Finish")
    return a + b + c
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res) // 1+2 = 3

    res = plusPlus(1, 2, 3) // Finish
		fmt.Println("1+2+3 =", res) // 1+2+3 = 6
}