
package main

import (
	"fmt"
)

func main() {
	var miles int
	var gallons int

	var totalMiles int
	var totalGallons int

	fmt.Print("Enter miles driven (-1 to quit): ")
	fmt.Scan(&miles)

	for miles != -1 {

		fmt.Print("Enter gallons used: ")
		fmt.Scan(&gallons)

		milesPerGallon := float64(miles) / float64(gallons)
		fmt.Println("Miles per gallon for this trip", milesPerGallon)

		totalMiles += miles
		totalGallons += gallons

		fmt.Print("Enter miles driven (-1 to quit): ")
		fmt.Scan(&miles)
	}

	if totalGallons != 0 {
		combinedMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Println("Combined miles per gallon:", combinedMPG)
	} else {
		fmt.Println("No trips were entered.")
	}
}