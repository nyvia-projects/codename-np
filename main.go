package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Loop to keep taking user input
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		// Convert the input string to lower case and trim spaces
		command := strings.TrimSpace(input)

		// Execute commands based on user input
		args := strings.Split(command, " ")
		switch args[0] {
		case "create":
			if len(args) != 2 {
				fmt.Println("Usage: create <filename.txt>")
			} else {
				createFile(args[1])
			}
		case "write":
			if len(args) < 3 {
				fmt.Println("Usage: write <filename.txt> <text>")
			} else {
				writeToFile(args[1], strings.Join(args[2:], " "))
			}
		case "read":
			if len(args) != 2 {
				fmt.Println("Usage: read <filename.txt>")
			} else {
				readFromFile(args[1])
			}
		case "ls":
			listFiles()
		case "ping":
			if len(args) < 2 {
				fmt.Println("Usage: ping <host>")
			} else {
				ping(args[1])
			}
		case "bye":
			fmt.Println("Goodbye!")
			return
		case "help":
			showHelp()
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}

//=============================================================================

// createFile creates an empty file with the given name
func createFile(filename string) {
	_, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File created:", filename)
}

// writeToFile writes text to a file
func writeToFile(filename, text string) {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Wrote to:", filename)
}

// readFromFile reads the content of a file and prints it
func readFromFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(content))
}

// listFiles lists all files and directories in the current directory
func listFiles() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func ping(host string) {
	cmd := exec.Command("ping", host)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print(out.String())
}

// showHelp prints a list of available commands
func showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  create <filename.txt>                - Creates a new file")
	fmt.Println("  write <filename.txt> <text>          - Writes text to a file")
	fmt.Println("  read <filename.txt>                  - Reads content from a file")
	fmt.Println("  ls                                   - Lists files in the current directory")
	fmt.Println("  ping <host>                          - Pings a host")
	fmt.Println("  bye                                  - Exits the program")
	fmt.Println("  help                                 - Shows this help")
}
