package main

import "fmt"

type arithmetic interface{ //SYNTAX : type interfaceName interface
	mathOperations() int //Function Signature
	//Function signature means the name of the function and its return type
}

//Implementing the interface using struct
type addition struct{
	number1,number2 int
}

type subtraction struct{
	number1,number2 int
}

type multiplication struct{
	number1,number2 int
}

func (a addition) mathOperations() int{
	sum := a.number1 + a.number2
	return sum
}

func (s subtraction) mathOperations() int{
	return s.number2 - s.number1
}

func (m multiplication) mathOperations() int{
	return m.number1 * m.number2
}

func result(arith arithmetic) int{
	return arith.mathOperations()
}

func main(){
	a := addition{number1 : 100, number2 : 50}
	b := subtraction{number1 : 30, number2 : 100}
	c := multiplication{number1 : 4, number2 : 50}
	fmt.Println(result(a))
	fmt.Println(result(b))
	fmt.Println(result(c))
}