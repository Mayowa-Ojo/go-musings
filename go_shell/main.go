package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // access input device

	for { // read input from keyboard
		workingDir, err := getWorkingDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s %v: ", green(string('\u203A')), green(workingDir))

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

	args := strings.Split(input, " ") // split args from main command

	// handle custom commands
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("invalid path")
		}

		return changeDir(args[1])
	case "exit":
		os.Exit(0)
	}

	// prepare command to be executed
	// pass the main command and rest of args
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run() // execute command

	return err
}

// --- custom shell commands
// change directory
func changeDir(dir string) error {
	err := os.Chdir(dir)

	return err
}

func getWorkingDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "nil", err
	}

	pathSlice := strings.Split(path, "/")
	dir := pathSlice[len(pathSlice)-1]

	return dir, nil
}
