package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	//context.Background() is used to create new context
	parentContext := context.Background()
	ctx, cancel := context.WithCancel(parentContext)
	//context.WithCancel(context.Background()) returns a context and a cancel function

	go runningProcess(ctx)
	time.Sleep(15*time.Second)

	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("Main Function Completed its Execution!")
}

func runningProcess(ctx context.Context){
	for {
		select{
		case <-ctx.Done(): //context pacakage will provide this. It will get executed when the cancel function gets its execution
			fmt.Println("Process cancelled")
			return
		default:
			fmt.Println("Running Process .... ")
			time.Sleep(1* time.Second)
		}
	}
}