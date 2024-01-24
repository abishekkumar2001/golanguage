package main

import (
	"fmt"
	"math/rand"
	"time"
)

func displayMessage(ch chan string) {
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("This is value %d in channel", getRand())
	}
	close(ch)
}

func getRand() int { //generating random number
	return rand.Intn(100)
}

func main() {

	//SYNTAX : channelName := make(chan return type,buffer)

	//Ignoring buffer value will result in a state known as buffer

	//Channel is a special variable that is used to exchange information about goroutines

	// ch<-"This is first message"
	// ch<-"This is second message"
	// ch<-"This is third message"
	// close(ch)

	//Close function to close the channel

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	ch := make(chan string, 3)
	go displayMessage(ch)
	for chvalue := range ch {
		fmt.Println(chvalue)
	}

}
