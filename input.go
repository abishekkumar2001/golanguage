package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Getting Integer as Input from user
	var n1, n2 int
	fmt.Print("Enter the value of n1 : ")
	fmt.Scan(&n1)
	fmt.Println("The value of n1 is ", n1)
	fmt.Print("Enter the value of n2 : ")
	fmt.Scanln(&n2)
	fmt.Println("The value of n2 is ", n2)

	//Getting Float as Input from user
	var a, b float32
	fmt.Print("Enter the value of a and b : ")
	fmt.Scanf("%v %v", &a, &b)
	fmt.Println("The value of a is", a, "and b is", b)
	fmt.Printf("The value of a is %v and b is %v\n", a, b)

	//Getting String as Input from user
	var name1,name2 string
	fmt.Println("Enter the name of first person :") 
	reader := bufio.NewReader(os.Stdin)
	name1, _ = reader.ReadString('\n')

	fmt.Println("Name of First person is", name1)

	// fmt.Scanln(&name1)   It will take only the first part of string

	fmt.Print("Enter the name of second person : ")
	reader = bufio.NewReader(os.Stdin)
	name2, _ = reader.ReadString('\n')

	fmt.Println("Name of Second person is",name2)

	//Getting Array as Input from user
	var n, i int
	fmt.Println("Enter the size of the array : ")
	fmt.Scan(&n)
	arr := make([]int, n) //creating an array or a slice using make function
	for i = 0; i < n; i++ {
		//	fmt.Printf("Enter the value of element at index %v",i)  Can be given like this also
		fmt.Printf("Enter the value of element at index %d", i)
		fmt.Scan(&arr[i])
	}

	//Printing the array elements
	fmt.Println("The elements of the array are")
	for i = 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
