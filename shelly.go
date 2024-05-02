package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	current_user, user_error := user.Current()
	if user_error != nil {
		os.Exit(0)
	}
	for {
		current_directory, directory_error := os.Getwd()
		if directory_error != nil {
			os.Exit(0)
		}
		fmt.Printf("(%s | %s) > ", current_user.Username, current_directory)
		//Read the keyboard input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		//Handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	//Remove the newline character
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	//Check in for built-in commands
	switch args[0] {
	case "cd":
		//'cd' to home dir with empty path not yet supported
		if len(args) < 2 {
			return errors.New("path required")
		}
		//cange the directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	//prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...)

	//Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//exec the command and return the error
	return cmd.Run()

}
