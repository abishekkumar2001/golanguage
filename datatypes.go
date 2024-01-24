package main
import "fmt"

func main(){
	// Integer
	fmt.Println("1.Integer")
	var a int =10
	var b=15
	c:=20
	fmt.Println("The value of a is ",a)
	fmt.Println("The value of b is ",b)
	fmt.Println("The value of c is ",c)
	fmt.Print("\n")  //for new line
	// Float
	fmt.Println("2.Float")
	var d=3.14
	fmt.Println("The value of d is ",d,"\n")
	//Boolean
	fmt.Println("3.Boolean")
	e:=true
	fmt.Printf("The value of e is %v and the type is %T",e,e)
	fmt.Println("\n")
	//String
	fmt.Println("4.String")
	var name string //variable declaration
	name="Steve" //variable Initialization
	fmt.Println("The name is",name)

}
