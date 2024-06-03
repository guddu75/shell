package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handlecommand(inputString string) {
	cmd, args := commandParser(inputString)
	switch cmd {
	case "exit":
		code, _ := strconv.Atoi(args[0])
		os.Exit(code)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "type":
		if args[0] == "exit" || args[0] == "echo" || args[0] == "type" {
			fmt.Println(args[0] + " is a shell builtin\n")
		} else {
			fmt.Println(args[0] + " not found\n")
		}
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

	for {

		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := reader.ReadString('\n')
		cmd = strings.Trim(cmd, "\n")
		if err != nil {
			log.Fatal(err.Error())
		}
		handlecommand(cmd)
	}

	// Wait for user input

}
