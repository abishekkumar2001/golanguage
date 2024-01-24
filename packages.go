package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	//Package fmt

	var a int
	fmt.Println("Enter the value of a : ")
	fmt.Scanln(&a)

	//Package math

	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Sqrt(27))
	fmt.Println(math.Cbrt(27))
	fmt.Println(math.Cbrt(25))
	fmt.Println(math.Min(27, 21))
	fmt.Println(math.Max(27, 21))

	//Package strings

	upper := strings.ToUpper("This is a Golang Program")
	fmt.Println(upper)
	lower := strings.ToLower("PACKAGES in GO")
	fmt.Println(lower)
}
