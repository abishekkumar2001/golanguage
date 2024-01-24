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
	wg.Add(5) 
	go myfunc1(&wg,5)
	go myfunc2(&wg,4)
	go myfunc1(&wg,3)
	go myfunc2(&wg,2)
	go myfunc1(&wg,1) //Address of waitgroup and id are passed as an argument
	wg.Wait() //Waits until the counter becomes zero
	fmt.Println("Main function Executed")
}

func myfunc1(wg *sync.WaitGroup,id int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Function1 Executed")
	time.Sleep(6 * time.Second)
	wg.Done() //We are confirming that the execution of this function gets completed
}


func myfunc2(wg *sync.WaitGroup,id int) {
	time.Sleep(12 * time.Second)
	fmt.Println("Function2 Executed")
	time.Sleep(6 * time.Second)
	wg.Done() //We are confirming that the execution of this function gets completed
}
