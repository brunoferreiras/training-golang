package main

import (
    "fmt"
    "time"
)

func main() {

    number := 2
    fmt.Print("Write ", number, " as ")
    switch number {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    now := time.Now()
    switch {
    case now.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(item interface{}) {
        switch typed := item.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", typed)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}