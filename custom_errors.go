package main

import "fmt"

type error interface{
	Error() string
}

type div struct{
	message string
}

func (d div) Error() string{ //Interface definition
	//SYNTAX : func (object structName) interfaceFunction() return type
	return "Mathematically a number cannot be divided by zero"
}

func division(n1 int,n2 int) (int,error){ //error is a data type in go
	if n2==0{
		return 0, &div{}
	} else {
		return n1/n2,nil
	}
}

func main(){
	var a,b int
	fmt.Println("Enter the value of a : ")
	fmt.Scanln(&a)

	fmt.Println("Enter the value of b : ")
	fmt.Scanln(&b)

	result,err := division(a,b)

	//Checking error occured or not
	if err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("The value is",result)
	}
}