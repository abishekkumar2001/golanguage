package main

import "fmt"

func main() {
	//Arithmetic operators
	var a, b int = 10, 5
	fmt.Println("The sum is", a+b)
	fmt.Println("The difference is", a-b)
	fmt.Println("The product is", a*b)
	fmt.Println("The quotient is", a/b) 
	fmt.Println("The remainder is", a%b)
	a++
	fmt.Println("The value after incrementing is", a)
	a--
	fmt.Println("The value after decrementing is", a)

	//Assignment operator
	var c = 20
	fmt.Println(c)

	//Comparison operators
	fmt.Println("Comparison operators")
	var i, j int = 10, 15
	fmt.Println(i > j)
	fmt.Println(i < j)
	fmt.Println(i == j)
	fmt.Println(i != j)
	fmt.Println(i >= j)
	fmt.Println(i <= j)

	//Logical operators
	fmt.Println("Logical operators")
	var n1, n2, n3 int = 100, 90, 170
	fmt.Println(n1 < n3 && n1 > n2) //AND operator
	fmt.Println(n1 > n3 && n1 > n2)
	fmt.Println(n1 > n3 || n1 < n2) //OR operator
	fmt.Println(n1 < n3 || n1 < n2)
	fmt.Println(!(n1 < n3 && n1 > n2)) //NOT operator
	fmt.Println(!(n1 > n3 || n1 < n2))

	//Bitwise operators
	var a1, a2 int = 4, 5
	fmt.Println(a1 & a2)
	fmt.Println(a1 | a2)
	fmt.Println(a1 ^ a2)

	//Left shift and right shift operators
	fmt.Println(a1 >> 2) //Right shift. Shift two places right
	fmt.Println(a1 << 2) //Left shift. Shift two places left
}
