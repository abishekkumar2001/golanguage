package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	parentContext := context.Background()
	ctx, cancel := context.WithTimeout(parentContext,3*time.Second)

	go timeoutProcess(ctx)
	time.Sleep(2*time.Second)

	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("Main function Executed!")
}

func timeoutProcess(ctx context.Context){  //Context message will be copied in ctx variable
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Process cancelled ")
			return
		default :
			fmt.Println("Process Running .....")
			time.Sleep(1*time.Second)
		}
	}
}