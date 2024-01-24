package main

import "fmt"

func main() {

	//Function declared Inside main

	marks := 65

	mark := func() int {
		marks = marks + 10
		return marks
	}
	fmt.Println("The value returned from the function is ", mark())
	fmt.Println(mark())
	fmt.Println(mark())

	//Function declared Outside main
	age1 := age()
	fmt.Println(age1())
	fmt.Println(age1())
	fmt.Println(age1())

	age2 := age()
	fmt.Println(age2())
	fmt.Println(age2())
	fmt.Println(age2())
}

func age() func() int {
	//SYNTAX : func functionName1() functionName2() return type
	//Here the function is returning another function of type int
	age := 90
	return func() int {
		age++
		return age
	}
}