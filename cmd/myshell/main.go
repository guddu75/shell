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

		// log.Printf("input", input)

		cmd, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err.Error())
		}

		cmd = strings.TrimSpace(cmd)

		fmt.Printf("%s: command not found", cmd)

		os.Exit(1)

	}

	// Wait for user input

}
