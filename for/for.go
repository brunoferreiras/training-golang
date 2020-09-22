package main

import "fmt"

func main() {
    number := 1
    for number <= 3 {
        fmt.Println(number)
        number = number + 1
    }

    for numberTwo := 7; numberTwo <= 9; numberTwo++ {
        fmt.Println(numberTwo)
    }

    for {
        fmt.Println("loop")
        break
    }

    for numberThree := 0; numberThree <= 5; numberThree++ {
        if numberThree%2 == 0 {
            continue
        }
        fmt.Println(numberThree)
    }
}