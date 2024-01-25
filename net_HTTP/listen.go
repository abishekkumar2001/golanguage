package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn){
	defer conn.Close()
	fmt.Println("Listened to ...",conn.RemoteAddr())
	//Handle the connection, read and write data
}

func main(){
	//Listen function from net package
	listener,err := net.Listen("tcp",":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening ... ")
	fmt.Println("Server is listening on port 8000")

	//Accepting incoming connections
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
			//Ignoring the error instance and continue listening other connections
		}
		go handleConnection(conn)
	}
}