package main

import "fmt"

func main(){

	//Multiple defer statements
	fmt.Println("one")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	defer fmt.Println("Four")  //The order of execution is LIFO
	
}