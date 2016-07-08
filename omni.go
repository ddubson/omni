package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify a command yo!")
		return
	}

	command := os.Args[1]

	if command == "commands" {
		fmt.Println("Supported commands:")
		fmt.Println("add-to-path")
	} else if command == "add-to-path" {
		fmt.Println("let's add something to the PATH variable")
	} else {
		fmt.Println("No such command. Run 'commands' to view what's available")
	}
}
