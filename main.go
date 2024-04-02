package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func main() {

	if len(os.Args) < 2 {
		// Show home menu
		seleted_option := home_menu()
		if seleted_option == "exit" {
			os.Exit(0)
		}
		if seleted_option == "send" {
			fmt.Print("Select a file to send\n")
			file := file_picker()
			if file == "" {
				fmt.Println("No file selected")
				os.Exit(1)
			}
			CreateServer(file)
		}
		if seleted_option == "receive" {
			var server_address string
			huh.NewInput().Title("Enter the server address").Value(&server_address).Run()
			if server_address == "" {
				fmt.Println("No server address provided")
				os.Exit(1)
			}
			// fmt.Printf("Receiving data from %s\n", server_address)
			createClient(server_address)
		}

		// CreateServer(file)
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
		// fmt.Println("Receiving data...")
		createClient(server_address)
		return
	}

}

func home_menu() string {

	var selected_option string

	huh.NewSelect[string]().Title("Choose an option").Options(
		huh.NewOption("Send File", "send"),
		huh.NewOption("Receive File", "receive"),
		huh.NewOption("Exit", "exit"),
	).Value(&selected_option).Run()

	return selected_option
}
