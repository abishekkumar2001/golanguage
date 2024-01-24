package main

import "fmt"

func main() {
	division(2,4)
	division(8,4)
	division(10,0)
	division(36,36)
}

func handlePanic(){

	//Detect if panic occurs or not
	c := recover()
	if c!=nil {
		fmt.Println("RECOVER",c)
	}
}

func division(a,b int){

	//Handle panic even after panic gets executed
	defer handlePanic()

	if b==0 {
		panic("Mathematical Error")
	} else {
		result := a/b
		fmt.Println("The value is ",result)
	}
	//The handlepanic function will get executed everytime but after the if ... else
	//If error occurs the recover function will get executed
}