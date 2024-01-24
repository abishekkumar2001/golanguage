package main

import "fmt"

func main(){
	fmt.Println("Hello World !!")
	fmt.Println("One")
	fmt.Println("TWO")
	fmt.Println("Three")
	panic("End of the program ... Statements after panic will not get executed ")
	//Statements after panic will not get executed
	fmt.Println("Four")
	fmt.Println("Five")
	fmt.Println("Six")
	fmt.Println("Seven")
	fmt.Println("Eight")
}
