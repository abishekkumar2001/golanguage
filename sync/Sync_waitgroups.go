package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Waitgroup from sync Package")
	//Waitgroup acts a counter. It will make sure that all functions get executed before the termination of main program
	var wg sync.WaitGroup //Waitgroup is defined in sync package
	wg.Add(2) 
	go myFunc(&wg,2) //Address of waitgroup is passed as an argument
	go myFunc(&wg,1)
	wg.Wait() //Waits until the counter becomes zero
	fmt.Println("Main function Executed")
}

func myFunc(wg *sync.WaitGroup, id int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Function1 Executed")
	time.Sleep(6 * time.Second)
	wg.Done() //We are confirming that the execution of this function gets completed
}
