package main

import "fmt"

func student() {
	var name string
	var grade string

	aCount := 0
	bCount := 0
	cCount := 0
	dCount := 0

	i := 0

	
	for i < 5 {
		fmt.Printf("Enter name of student %d: ", i+1)
		fmt.Scan(&name)

		fmt.Printf("Enter grade for %s (A-D): ", name)
		fmt.Scan(&grade)

		
		switch grade {
		case "A", "a":
			aCount++
		case "B", "b":
			bCount++
		case "C", "c":
			cCount++
		case "D", "d":
			dCount++
		default:
			fmt.Println("Invalid grade. Please enter A, B, C, or D.")
			continue 
		}

		i++
		fmt.Println()
	}

	
	fmt.Println("Grade Summary:")
	fmt.Println("A:", aCount)
	fmt.Println("B:", bCount)
	fmt.Println("C:", cCount)
	fmt.Println("D:", dCount)
}
