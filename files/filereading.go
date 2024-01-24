package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	readFile()
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
