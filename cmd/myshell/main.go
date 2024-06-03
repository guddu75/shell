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
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := reader.ReadString('\n')
		cmd = strings.Trim(cmd, "\n")

		if err != nil {
			log.Fatal(err.Error())
		}

		handlecommand(cmd)

		// fmt.Println(cmd)

		// if cmd == "exit 0" {
		// 	os.Exit(0)
		// } else if strings.HasPrefix(cmd, "echo") {
		// 	fmt.Println(cmd[5:])
		// } else {
		// 	fmt.Printf("%s: command not found\n", cmd)
		// }
		// os.Exit(0)

		// os.Exit(1)

	}

	// Wait for user input

}
