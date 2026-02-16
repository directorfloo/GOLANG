package main

import "fmt"

func calculator() {
	for {
		var accountNumber int

		fmt.Print("Enter account number (-1 to quit): ")
		fmt.Scan(&accountNumber)

		if accountNumber == -1 {
			fmt.Println("Exiting program...")
			break
		}

		var beginningBalance, charges, credits, creditLimit float64

		fmt.Print("Enter beginning balance: ")
		fmt.Scan(&beginningBalance)

		fmt.Print("Enter total charges: ")
		fmt.Scan(&charges)

		fmt.Print("Enter total credits: ")
		fmt.Scan(&credits)

		fmt.Print("Enter credit limit: ")
		fmt.Scan(&creditLimit)

		newBalance := beginningBalance + charges - credits

		fmt.Println("\nAccount Number:", accountNumber)
		fmt.Println("Credit Limit:", creditLimit)
		fmt.Println("New Balance:", newBalance)

		if newBalance > creditLimit {
			fmt.Println("Status: Credit limit exceeded!")
		} else {
			fmt.Println("Status: Within credit limit.")
		}
		fmt.Println()
	}
}
