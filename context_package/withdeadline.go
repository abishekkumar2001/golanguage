package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	parentContext := context.Background()
	deadline := time.Now().Add(3*time.Second)
	ctx, cancel := context.WithDeadline(parentContext,deadline)

	go deadlineProcess(ctx)
	time.Sleep(2*time.Second)

	defer cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("Main function Executed!")
}

func deadlineProcess(ctx context.Context){  //Context message will be copied in ctx variable
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Process cancelled ")
			return
		default :
			fmt.Println("Running .....")
			time.Sleep(1*time.Second)
		}
	}
}