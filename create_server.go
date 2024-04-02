package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func CreateServer(file_path string) {
	// Start TCP server on port 8080
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server started. Listening on port 8080...")

	for {
		// Accept incoming connections
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle connection in a separate goroutine
		go handleConnection(conn, file_path)
	}
}

func handleConnection(conn net.Conn, file_path string) {
	defer conn.Close()

	// Open the file
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	filename := filepath.Base(file_path)
	fmt.Println("Filename:", filename)
	_, err = conn.Write([]byte(filename))
	if err != nil {
		fmt.Println("Error sending filename:", err)
		return
	}

	// Copy file contents to the client connection
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file contents:", err)
		return
	}
}
