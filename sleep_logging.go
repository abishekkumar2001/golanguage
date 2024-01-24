package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	//Calling sleep method

	time.Sleep(8 * time.Second)

	fmt.Println("Sleep Function Executed")
	time.Sleep(3 * time.Second)
	fmt.Println("Executed...")

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	//Logger writing to standard output with prefix info

	logger.Println("This is an info message.")
	//Printing messages using logger
	
	logger.Fatal("Fatal Error")
	//Printing error messages using logger
}
