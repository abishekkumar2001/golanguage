package main

import "fmt"

//Functions in go

//Function declaration

func add(){ //Function with no parameters
	var a,b,sum int
	fmt.Println("Enter the Value of a : ")
	fmt.Scanln(&a)
	fmt.Println("Enter the Value of b : ")
	fmt.Scanln(&b)
	sum = a+b
	fmt.Println("The sum is ",sum)
}

//Function with parameters

func arithmeticOperations(a int,b int){
	var c int
	c=a+b
	fmt.Printf("The addition of two numbers %d and %d is %d \n",a,b,c)
	c=a*b
	fmt.Printf("The multiplication of two numbers %d and %d is %d \n",a,b,c)
}

func main(){
	add() //Function call
	arithmeticOperations(10,20)
	var number int
	fmt.Println("Enter the value of number : ")
	fmt.Scanf("%d",&number)
	var n int = countDigits(number)
	fmt.Println("The number of digits in the given number is",n)
	result := cube(6)
	fmt.Println("The cube of the number is", result)
	var sqNumber int 
	fmt.Println("Enter the value : ")
	fmt.Scanf("%d\n",&sqNumber)
	fmt.Println(square(sqNumber))
}

//Function with return values
func countDigits(num int)int {
	var count int = 0
	for num!=0 {
		num=num/10
		count++
	}
	return count
}

//Function with named return value
func cube(num int) (cube int) {
	cube =num*num*num
	//cube := num*num*num
	return
}

//Function with multiple return values
func square(num int) (square int, squareRoot int){
	square = num*num
	var i int =1
	var result,flag int
	flag=0
	for i=1; i<(num)/2; i++{
		result = i*i
		if(result == num){
			squareRoot = i
			flag=1
			break
		}
	}
	if(flag == 0){
		fmt.Println("It is not a perfect square")
	}
	return square,squareRoot
}
