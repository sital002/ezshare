package main

import (
	"errors"
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
		if opErr, ok := err.(*net.OpError); ok && opErr.Op == "listen" {
			fmt.Println("The port is already in use")
		} else {
			fmt.Println("Error occurred while trying to listen on the port:", err)
		}
		return
	}
	hostIP, err := getIPFromInterface("Wi-Fi")
	if err != nil {
		fmt.Println("Error getting host IP address:", err)
		return
	}

	fmt.Println("Host IP address:", hostIP)

	fmt.Println("Server started. Listening on port 8080...")
	// ...

	defer ln.Close()

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

func getIPFromInterface(interfaceName string) (string, error) {
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()
		if ip == nil {
			continue
		}

		return ip.String(), nil
	}

	return "", errors.New("no suitable IP address found")
}
