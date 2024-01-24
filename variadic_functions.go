package main

import "fmt"

//Variadic function

//Can be used to pass variable number of parameters. It behaves like a slice 

//Variadic function SYNTAX : func functionName(parameter...type)return type

func sum(val ...int) int {
	total := 0
	//For Range loop SYNTAX : for index, value := range arrayName
	//Here arrayName is val... result is the variable we have declared for the elements of val
	for _, result := range val {
		total = total + result
	}
	return total
}

func main() {

	//Function call
	fmt.Println(sum(10, 12, 15, 12, 18)) //Sum function is called with five parameters

	a := sum(20, 30, 50)
	fmt.Println("The sum is ",a) //Sum function is called with three parameters

	fmt.Println(sum(1020, 23)) //Sum function with two parameters

	//Variadic function can be applied to slices
	nums := []int{1,2,3,4,5}
	fmt.Println(sum(nums...))
}
