package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	
	//Connecting to the TCP Server on localhost
	conn,err := net.Dial("tcp","localhost:8080")
	if err != nil {
		fmt.Println("Error Connecting",err)
		os.Exit(1)
	}
	defer conn.Close()

	//Sending message to the server
	msg := "Hello Server !"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error Writing to server",err)
		return
	}
	fmt.Printf("Message Sent is %s",msg)

	//Echoed message from server
	buffer := make([]byte,1024)
	n,err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error Reading",err)
		return
	}
	fmt.Printf("\nEchoed message is %s\n",buffer[:n])

}