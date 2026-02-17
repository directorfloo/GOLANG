package main

import "fmt"

func christmas() {

	for day := 1; day <= 12; day++ {

		var dayName string

		switch day {
		case 1:
			dayName = "first"
		case 2:
			dayName = "second"
		case 3:
			dayName = "third"
		case 4:
			dayName = "fourth"
		case 5:
			dayName = "fifth"
		case 6:
			dayName = "sixth"
		case 7:
			dayName = "seventh"
		case 8:
			dayName = "eighth"
		case 9:
			dayName = "ninth"
		case 10:
			dayName = "tenth"
		case 11:
			dayName = "eleventh"
		case 12:
			dayName = "twelfth"
		}

		fmt.Printf("\nOn the %s day of Christmas my true love sent to me:\n", dayName)

	
		switch day {

		case 12:
			fmt.Println("Twelve drummers drumming,")
			fallthrough
		case 11:
			fmt.Println("Eleven pipers piping,")
			fallthrough
		case 10:
			fmt.Println("Ten lords a-leaping,")
			fallthrough
		case 9:
			fmt.Println("Nine ladies dancing,")
			fallthrough
		case 8:
			fmt.Println("Eight maids a-milking,")
			fallthrough
		case 7:
			fmt.Println("Seven swans a-swimming,")
			fallthrough
		case 6:
			fmt.Println("Six geese a-laying,")
			fallthrough
		case 5:
			fmt.Println("Five golden rings,")
			fallthrough
		case 4:
			fmt.Println("Four calling birds,")
			fallthrough
		case 3:
			fmt.Println("Three French hens,")
			fallthrough
		case 2:
			fmt.Println("Two turtle doves,")
			fallthrough
		case 1:
			if day == 1 {
				fmt.Println("A partridge in a pear tree.")
			} else {
				fmt.Println("And a partridge in a pear tree.")
			}
		}
	}
}
