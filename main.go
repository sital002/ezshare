package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		// fmt.Println("No arguments provided")
		file := file_picker()
		if file == "" {
			fmt.Println("No file selected")
			return
		}
		CreateServer(file)
		return
	}
	if os.Args[1] == "send" {
		if os.Args[2] == "" {
			fmt.Println("No file provided")
			return
		}
		file_path := os.Args[2]
		CreateServer(file_path)
		return
	}
	if os.Args[1] == "receive" {
		if os.Args[2] == "" {
			fmt.Println("No server address provided")
			return
		}
		server_address := os.Args[2]
		fmt.Println("Receiving data...")
		createClient(server_address)
		return
	}

}
