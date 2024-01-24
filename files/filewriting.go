package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	createFile()
	readFile()
}

func createFile() {
	fmt.Printf("Writing to a file in Go lang\n")

	//Creating a text file
	file, err := os.Create("file.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		//SYNTAX : log.fatal("Message : %s",err)
	}

	defer file.Close()

	len, err := file.WriteString("Hello" + " This program demonstrates reading and writing" + " operations to a file in Go lang.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func readFile() {
	fmt.Println("File reading Demo")
	fileName := `file.txt`

	//For reading a file... It should be created earlier

	//Reding the contents of file
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Panicf("Unable to read file %s", err)
		//SYNTAX : log.panicf("Message %s",err)
	}

	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s\n", data)
}
