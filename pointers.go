package main

import "fmt"

func main() {

	var a int //Normal Integer Variable
	fmt.Println("Enter the value of a : ")
	fmt.Scanln(&a)

	var ptr *int //Integer Pointer variable
	ptr = &a     //Ptr variable is a variable that holds the address of another variable.

	//Here ptr is holding the address of a

	fmt.Println("The value of a is ", a)
	fmt.Println("The value of ptr is ", ptr)

	//To change the value in the address which ptr is holding.
	*ptr = a + 5

	fmt.Println("The value is ", a)
	//The value stored in a will get printed.
	fmt.Println("The contents of ptr is", *ptr)
	//The value stored by a will get printed. Since ptr is pointing the adress of a.

	var b int
	fmt.Println("Enter the value of b : ")
	fmt.Scanln(&b)

	fmt.Println("The values before swaping...")
	fmt.Println("The value of a is : ", a)
	fmt.Println("The value of b is : ", b)
	swap(&a, &b)
	fmt.Println("The values after swaping...")
	fmt.Println("The value of a is : ", a)
	fmt.Println("The value of b is : ", b)
}

//we can also write like func swap(a1,b1 *int)
func swap(a1 *int, b1 *int) { // func swap(argument 1, argument 2)(return type)
	*a1, *b1 = *b1, *a1
}
