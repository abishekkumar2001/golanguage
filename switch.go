package main

import "fmt"

func main() {

	var month int
	fmt.Println("Enter the number for which you need month")
	fmt.Scanln(&month)
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Not a valid month")
	}

	//fallthrough keyword
	var num int
	fmt.Println("Enter the number")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("Extraordinary")
		fallthrough
	case 2:
		fmt.Println("Outstanding")
	case 3:
		fmt.Println("Excellent")
	case 4:
		fmt.Println("Very Good")
	case 5:
		fmt.Println("Good")	
	default:
		fmt.Println("TNot Good")
	}
}
