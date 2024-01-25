package main

import (
	"fmt"
	"net"
)

func main(){
	conn,err := net.Dial("tcp","example.com:80")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to the server ....... ")
}