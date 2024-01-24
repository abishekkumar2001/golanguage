package main

import (
	"bufio"
	"fmt"
	"os"
)

type employee struct { //structure 1
	name   string
	age    int
	salary int
	job    string
}

type school struct { //structure 2
	name     string
	class    int
	marks    int
	comments string
}

//Nested structure
/*
	SYNTAX : 

	type structName1 struct{
		field1
		field2
		......
	}

	type structName2 struct{
		variableName structName1
		field1
		field2
		......
	}		 
*/

type personal struct{
	maritalStatus string
	vehicles int
	details employee //nested strcuct
	marks int

}

func main() {
	var emp1 employee
	var emp2 employee

	//Details of first employee
	emp1.name = "Sachin Tendulkar"
	emp1.age = 25
	emp1.salary = 50000
	emp1.job = "Junior Software Engineer"

	/*Can also be defined like
	  
	  var emp1 employee
	  var emp1 = employee{"Sachin Tendulkar",25,50000,"Junior Software Engineer"}

	  var emp2 employee
	  var emp2 = employee{name = "Sachin Tendulkar",
	   age = 25,
	   salary = 50000,
	   job = "Junior Software Engineer"
	 }
	*/

	//Details of second employee
	emp2.name = "Rohit Sharma"
	emp2.age = 32
	emp2.salary = 120000
	emp2.job = "Senior Software Engineer"

	//Printing the details
	fmt.Println(emp1) //Prints the entire details of employee1
	fmt.Println("Designation of second employee is ", emp2.job)

	//Getting Input from User
	var student school
	fmt.Println("Enter the name of student : ")
	reader := bufio.NewReader(os.Stdin)
	student.name, _ = reader.ReadString('\n')

	fmt.Println("Which class he is studying currently? ")
	fmt.Scanf("%d\n", &student.class)
	fmt.Println("Enter the marks he obtined out of 500 : ")
	fmt.Scanf("%d\n", &student.marks)
	fmt.Println("Your comment : ")
	fmt.Scanf("%s\n", &student.comments)

	//Printing the details
	fmt.Println("The details of student : ", student)
	fmt.Println("The comments given is ", student.comments)

	//nested structure
	//Initializing the function fields
	record := personal {
		maritalStatus : "Unmarried",
		vehicles : 5,
		details : employee{
			name :"Virat",
			age : 32,
			salary : 200000,
			job : "Cricketer",
		},
		marks : 1100,
	}

	//Display the contents of nested structure
	fmt.Printf("%v",record)
}
