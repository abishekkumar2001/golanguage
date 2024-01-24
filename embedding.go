package main

import (
	"fmt"
	"math"
)

//Definitions for embedding Structure

type employee struct {
	name   string
	salary int
	exp    int
	job    string
}

type personal struct {
	gender  string
	age     int
	details employee
	marks   int
}

func main() {

	//Embedding Structure
	var emp employee

	emp.name = "Akshay"
	emp.salary = 200000
	emp.exp = 15
	emp.job = "Senior Software Engineer"

	//Nested structure

	display := personal{
		gender: "Male",
		age:    35,
		details: employee{
			name:   "Rohit Sharma",
			salary: 300000,
			exp:    5,
			job:    "Junior Software Engineer",
		},
		marks: 1189,
	}

	fmt.Println(display)
	fmt.Println("The Job given is ", display.details.job)

	//Embedding Interface

	//Assigning values to the sructure

	values := mathematics{
		number1: 50,
		number2: 100,
		radius:  3.5,
		breadth: 4.5,
		height:  10,
	}

	var f finalInterface = values
	fmt.Println("The addition value is", f.mathOperations())
	fmt.Println("The area of circle is", f.area())
	fmt.Println("The area of rectangle is", f.area1())

}

//Definitions for embedding Interface

type arithmetic interface {
	mathOperations() int
}

type shapes interface {
	area() float64
}

type shapes1 interface{
	area1() float64
}

type finalInterface interface {
	arithmetic //Interface name1
	shapes  //Interface name2
	shapes1 //Interface name3
}

type mathematics struct {
	number1 int
	number2 int
	radius  float64
	breadth float64
	height  float64
}

// Implementing method in interface 1
func (m mathematics) mathOperations() int {
	return m.number1 + m.number2
}

// Implementing method in interface 2
func (m mathematics) area() float64 {
	return math.Pi * m.radius * m.radius
}

func (m mathematics) area1() float64{
	return m.breadth * m.height
}