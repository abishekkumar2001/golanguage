package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name1, name2 string

	//Getting Input from user
	fmt.Println("Enter the name of the first person :")
	reader := bufio.NewReader(os.Stdin)
	name1, _ = reader.ReadString('\n')

	fmt.Println("Name of First person is", name1)

	// fmt.Scanln(&name1)   It will take only the first part of string

	fmt.Print("Enter the name of the second person : ")
	reader = bufio.NewReader(os.Stdin)
	name2, _ = reader.ReadString('\n')

	fmt.Println("Name of Second person is", name2)

	//Representing string using back tick

	var str = `Strings in go programming`
	fmt.Println("The string is", str)

	//Accessing characters from string

	fmt.Printf("%c\n", name1[2])
	fmt.Printf("%c\n", str[2])

	//Counting the length of the string
	var strlen int
	strlen = len(str)
	fmt.Println("The length of the string is", strlen)

	//String operations
	//String concatenation

	str1 := "His favourite"
	str2 := "food is fried rice"
	str3 := str1 + str2
	fmt.Println(str3)

	//String compare
	var person1, person2, person3 string = "Sachin", "Rohit Sharma", "Sachin"
	fmt.Println(strings.Compare(person1, person2)) // Return 1 since, string 1 is smaller than string 2
	fmt.Println(strings.Compare(person2, person3)) // Returns -1 since, string 2 is greater than string 3
	fmt.Println(strings.Compare(person1, person3)) //Returns 0 since, both are equal

	//Substring checking using string contains function
	str4 := "is"
	result := strings.Contains(str1,str4)
	fmt.Println(result)

	//String replace function
	replace := strings.Replace(str1,"His","Everyone's",3)
	fmt.Println(replace)

	//Changing the case
	res1 := strings.ToUpper(str2)
	res2 := strings.ToLower(str1)

	fmt.Println(res1)
	fmt.Println(res2)
}


