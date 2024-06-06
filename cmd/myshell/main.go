package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var paths []string

func exitCmd(arg string) {
	code, _ := strconv.Atoi(arg)
	os.Exit(code)
}

func echoCmd(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func typeCmd(args []string) {
	if args[0] == "exit" || args[0] == "echo" || args[0] == "type" {
		fmt.Println(args[0] + " is a shell builtin")
	} else {
		for _, path := range paths {
			filepath := filepath.Join(path, args[0])
			if _, err := os.Stat(filepath); err == nil {
				fmt.Println(args[0] + " is " + filepath)
				return
			}
		}
		fmt.Printf("%s: command not found\n", args[0])
		// os.Exit(0)

	}
}

func handlecommand(inputString string) {
	cmd, args := commandParser(inputString)
	switch cmd {
	case "exit":
		exitCmd(args[0])
	case "echo":
		echoCmd(args)
	case "type":
		typeCmd(args)
	default:
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func commandParser(cmd string) (string, []string) {
	tokens := strings.Split(cmd, " ")

	if len(tokens) < 2 {
		return tokens[0], []string{}
	}

	return tokens[0], tokens[1:]

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	path := os.Getenv("PATH")

	paths = strings.Split(path, ":")

	for {

		fmt.Printf("$ ")
		cmd, err := reader.ReadString('\n')
		cmd = strings.Trim(cmd, "\n")
		if err != nil {
			log.Fatal(err.Error())
		}
		handlecommand(cmd)
	}

	// Wait for user input

}
