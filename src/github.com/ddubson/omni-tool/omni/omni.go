package main

import (
	"fmt"
	"os"
	"os/user"
	"os/exec"
	"log"
	"github.com/ddubson/omni-tool/omnilib"
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

		dirToAdd := os.Args[2]

		addToLocalBashProfile(dirToAdd)
	} else {
		fmt.Println("No such command. Run 'commands' to view what's available")
	}
}

func addToLocalBashProfile(directory string) {
	bash := getUserHomeDir() + "/" + omnilib.BashProfile
	pathString := fmt.Sprintf(omnilib.PathCommand, directory)
	cmdString := fmt.Sprintf(omnilib.ExportCommand + " " + pathString)
	//fmt.Println(cmdString)

	f, err := os.OpenFile(bash, os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	_, err2 := f.WriteString(cmdString)

	if err2 != nil {
		fmt.Fprintf(os.Stderr, err2.Error())
	}

	f.Close()
	out, _ := exec.Command("source", bash).Output()
	fmt.Printf("%s", out)

	out2, _ := exec.Command("export", pathString).Output()
	fmt.Printf("%s", out2)

	fmt.Printf("### Success! ### \nYour directory has been added to the local Path.\nOpen a new terminal to view changes.\n")
}

func getUserHomeDir() (string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
