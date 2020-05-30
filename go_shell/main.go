package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // access input device

	for { // read input from keyboard
		fmt.Println("-> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = executeInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// executes command from keyboard input
func executeInput(input string) error {
	input = strings.TrimSuffix(input, "\n") // strip newline character

	cmd := exec.Command(input) // prepare command to be executed

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run() // execute command

	return err
}
