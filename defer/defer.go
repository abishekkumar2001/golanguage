package main

import "fmt"

func main(){

	//Defer statement ... It is used to delay the execution of a particular statement
	fmt.Println("one")
	defer fmt.Println("Two")
	fmt.Println("Three")
	fmt.Println("Four")

}