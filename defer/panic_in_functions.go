package main

import "fmt"

func main(){
	//Panic Example
	divide(8,2)
	divide(49,7)
	divide(10,0)
	//Satements after this will not get executed
	divide(9,3)
	divide(39,3)
	divide(100,5)
}

func divide(a,b int) { //Function with no return type
	if b == 0{
		panic("Mathematical error")
	} else {
		result := a/b
		fmt.Println("The value is ",result)
	}
}