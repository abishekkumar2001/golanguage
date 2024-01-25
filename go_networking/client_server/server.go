package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Read data from the connection
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		data := buffer[:n]
		fmt.Printf("Received data: %s\n", data)

		// Echo the data back to the client
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func main() {
	// Start a TCP server on port 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on localhost:8080")

	for {
		// Accept a connection from a client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		
		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}
