// Package declaration: Use lowercase and short names for packages.
package main

import (
	"fmt"
	"net/http"
)

// Constants: Use uppercase for exported constants.
const BaseURL = "https://example.com/api"

// Struct: Use CamelCase for struct names and their fields.
type Person struct {
	FirstName string
	LastName  string
}

// Function: Use CamelCase for function names.
func main() {
	// Variable declaration: Use camelCase for variable names.
	var greetingMessage string //We can also use Pascal/Snake case for naming a variable
	greetingMessage = "Hello, World!"

	// Print the greeting message.
	fmt.Println(greetingMessage)

	// Call a function with a receiver (method).
	p := Person{FirstName: "Akshay", LastName: "Kumar"}
	p.PrintFullName()

	// Example of using a third-party package (net/http).
	resp, err := http.Get(BaseURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}

// Method: Use CamelCase for method names.
func (p Person) PrintFullName() {
	fmt.Printf("Full Name: %s %s\n", p.FirstName, p.LastName)
}
