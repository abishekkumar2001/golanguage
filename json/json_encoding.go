package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string //Note the first letter is to be given in CAPITAL letter
	Age    int
	Salary int
	Gender string
}

func main() {

	person1 := Person{Name: "Akshay", Age: 60, Salary: 10000, Gender: "Male"}

	//Encoding person1 struct into json format

	personEnc, err := json.Marshal(person1)
	if err != nil {
		fmt.Println(err)
	}

	//As personEnc is a byte array it needs to be converted into a string

	fmt.Println(string(personEnc))

	//Converting slices into json format
	person2 := []Person{
		{Name : "Sachin", Age : 27, Salary : 15000, Gender : "Male"},
		{Name : "Sanjana", Age : 25, Salary : 13000, Gender : "Female"},
		{Name : "Rithu", Age : 28, Salary : 17000, Gender : "Female"},
		{Name : "Aravind", Age : 23, Salary : 13000, Gender : "Male"},
	}

	perEnc, err := json.Marshal(person2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The contents of person 2 are")
	fmt.Println(person2)
	fmt.Println("JSON enoding")
	fmt.Println(perEnc) //Will display in byte format

	//As perEnc is a byte array it needs to be converted into a string
	fmt.Println("JSON Array in string format")
	fmt.Println(string(perEnc))
}
