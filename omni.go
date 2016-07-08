package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"os/user"
	"log"
)

const bashProfile string = ".bash_profile"

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
		printLocalBashProfile()
	} else {
		fmt.Println("No such command. Run 'commands' to view what's available")
	}
}

func printLocalBashProfile() {
	// read the whole file at once
	file, err := ioutil.ReadFile(getUserHomeDir() + "/" + bashProfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	} else {
		fmt.Println(string(file))
	}
}

func getUserHomeDir() (string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
