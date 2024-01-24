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
	var person1 Person
	//Structure variable

	//Data in json format which is to be decoded
	Data := []byte(`{
		"Name" : "Sanjay",
		"Age" : 50,
		"Salary" : 10000,
		"Gender" : "male"
	}`)
	fmt.Println(Data) //The encoded value will get printed

	//Decoding person1 from json format
	err := json.Unmarshal(Data, &person1)
	if err != nil { 
		fmt.Println(err) 
    } 
	
	//Displaying the details of decoded data
	fmt.Println("The structure values are")
	fmt.Println(person1)
	fmt.Println("The age is",person1.Age)

	var person2 []Person
	//Array declaration for the structure person

	Data1 := []byte(`[
		{"Name" : "Sachin", "Age" :25, "Salary" : 100000, "Gender" : "Male"},
		{"Name" : "Rohit", "Age" :27, "Salary" : 200000, "Gender" : "Male"},
		{"Name" : "Sweety", "Age" :23, "Salary" : 500000, "Gender" : "Female"}	
	]`)
	fmt.Println(Data1)

	err1 := json.Unmarshal(Data1,&person2)
	if err1 != nil { 
		fmt.Println(err1) 
    }
	
	fmt.Println(person2)

	//Using for loop
	for i := range person2{
		fmt.Println(person2[i])
	}
}
