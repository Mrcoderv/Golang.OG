package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer listener.Close()

	fmt.Println("TCP Server running on port 8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		fmt.Println("Client connected")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	conn.Write([]byte("Hello from TCP Server!\n"))

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Client says:", string(buffer[:n]))
}