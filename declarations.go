package main
import "fmt"

var value=100

func main(){

	//Using var keyword and datatype
	fmt.Println("Using var keyword and type")
	var a int =5
	fmt.Println("The value of a is ",a)
	//Using var keyword and without datatype
	fmt.Println("Using var keyword and without type")
	var b=true
	fmt.Println(b)
	//Using := sign
	fmt.Println("Using := sign")
	c:=55.6
	fmt.Println("The value of c is",c)
	//Without initial value
	var name string
	name="Sanjay"
	fmt.Println(name)
	//Value declared outside the function
	fmt.Println("Value declared outside the function is",value)
	//multiple declarations
	var fruit1,fruit2,fruit3 string="Apple","Orange","Grapes"
	fmt.Println(fruit1)
        fmt.Println(fruit2)
	fmt.Println(fruit3)
}
