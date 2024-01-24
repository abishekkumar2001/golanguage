package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Enter the value of a : ")
	fmt.Scanln(&a)
	fmt.Println("Enter the value of b :")
	fmt.Scanln(&b)
	result := divide(a, b)
	fmt.Println("The value is ", result)

	//Error handling using errorf() function

	var time int
	fmt.Println("Enter the value of time : ")
	fmt.Scanf("%d", &time)

	//Creating error using errorf() function

	error := fmt.Errorf("%d is negative \nTime cannot be negative ", time)
	if time < 0 {
		fmt.Println(error)
	} else {
		fmt.Println("The time Entered is ", time)
	}


	err := division(a,b)

	if err != nil{
		fmt.Printf("error : %s",err)
	} else {
		fmt.Println("Valid division")
	}
 
}

//Error handling using new function

func divide(m int, n int) int {
	var result int
	if n == 0 {
		mathError := errors.New("Mathematical error")
		fmt.Println(mathError)
	} else {
		result = m / n
		return result
	}
	return result
}

//Function returning error

func division(a,b int) error {

	//Returns error
	if b == 0 {
		return fmt.Errorf("Mathematically a value cannot be divided by zero ")
	}

	//Returns zero
	return nil
}