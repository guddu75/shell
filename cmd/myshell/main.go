package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if err != nil {
			log.Fatal(err.Error())
		}

		if cmd == "exit 0" {
			os.Exit(0)
		} else if strings.HasPrefix(cmd, "$ echo") {
			fmt.Printf(cmd[7:] + "\n")
		}

		fmt.Printf("%s: command not found\n", cmd)

		// os.Exit(1)

	}

	// Wait for user input

}
