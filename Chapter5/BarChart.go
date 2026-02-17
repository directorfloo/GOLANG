package main

import "fmt"

func chart() {
	var numbers [5]int
	for i := 0; i < 5; i++ {
		for {
			fmt.Printf("Enter number %d (1-30): ", i+1)
			fmt.Scan(&numbers[i])

			if numbers[i] >= 1 && numbers[i] <= 30 {
				break
			}
			fmt.Println("Invalid input. Please enter a number between 1 and 30.")
		}
	}

	fmt.Println("\nBar Chart:")

	for _, num := range numbers {
		for i := 0; i < num; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
