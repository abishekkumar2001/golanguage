package main

import (
	"goProject/greeting"
	"goProject/authentication"
	"fmt"
	"os"
)

// make subTotalBill as global variable to make it easily accessible in case customer modifies her order.
// subTotalBill means total bill but excluding taxes.
var subTotalBill float64

// make a map of customerOrder in which we will store the items ordered as "key" and no. of plates as "value".
var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greeting.Greet(customerName)

	var loggedIn bool

	for !loggedIn {
		fmt.Println("1. Signup")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			signup()
		case 2:
			if err := login(); err != nil {
				fmt.Println("Login unsuccessful. Exiting.")
				os.Exit(1)
			}
			loggedIn = true
		case 3:
			fmt.Println("Exiting.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}


	orderItems()
	displayGeneratingBill() //just displays that "generating bill" in a fancy manner.
	generateBill()
	modifyOrder()
	printFinalBill()
	greeting.SayTata(customerName) //SayTata is defined in greeting package
}

func signup() {
	var username, password string

	fmt.Print("Enter your desired username: ")
	fmt.Scan(&username)

	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	// Perform signup
	response, err := authentication.SignupUser(username, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(response)
}

func login() error {
	var username, password string

	fmt.Print("Enter your username: ")
	fmt.Scan(&username)

	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	// Perform login
	response, err := authentication.LoginUser(username, password)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println(response)

	// If login is unsuccessful, return an error
	if response == "Invalid username or password" {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}


