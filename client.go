package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func createClient(server_address string) {
	// Connect to TCP server on localhost:8080
	conn, err := net.Dial("tcp", server_address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Create a buffer to store received data

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving filename:", err)
		return
	}
	filename := string(buffer[:n])

	// Create a file to save received data
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Receive data from the server
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Error receiving file contents:", err)
		return
	}

	fmt.Printf("File '%s' received and saved.\n", filename)
}
