package main

import "fmt"

func meters() {
    var feet float64
    var meters float64

    fmt.Print("Enter value in feet: ")
    fmt.Scanln(&feet)

    meters = feet * 0.3048

    fmt.Println("Value in meters =", meters)
}
