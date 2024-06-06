package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var paths []string

var builtinCommands = make(map[string]int)

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
		fmt.Println(args[0] + ": not found")
		// fmt.Printf("%s: not found", args[0])
		// os.Exit(0)

	}
}

func execCmd(file string, args []string) {
	for _, path := range paths {
		filepath := filepath.Join(path, file)
		if _, err := os.Stat(filepath); err == nil {
			cmd := exec.Command(filepath, args...)
			ouput, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("%s", string(ouput))
			return
		}
	}
	fmt.Printf("%s: command not found\n", file)
	return
}

func isbuiltIn(cmd string) bool {
	_, ok := builtinCommands[cmd]
	return ok
}

func pwdCmd() {
	currDir, err := Pwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(currDir)
}

func cdCmd(path string) {
	log.Fatal(path)
	err := Cd(path)
	if err != nil {
		fmt.Println("cd: " + path + ": No such file or directory")
	}
}

func handlecommand(inputString string) {
	cmd, args := commandParser(inputString)

	if cmd == "exit" {
		exitCmd(args[0])
	} else if cmd == "echo" {
		echoCmd(args)
	} else if cmd == "type" {
		typeCmd(args)
	} else if !isbuiltIn(cmd) {
		execCmd(cmd, args)
	} else if cmd == "pwd" {
		pwdCmd()
	} else if cmd == "cd" {
		cdCmd(args[0])
	} else {
		fmt.Printf("%s: command not found\n", cmd)
	}

	// switch cmd {
	// case "exit":
	// 	exitCmd(args[0])
	// case "echo":
	// 	echoCmd(args)
	// case "type":
	// 	typeCmd(args)
	// case isbuiltIn():
	// 	execCmd(args)
	// default:

	// }
}

func commandParser(cmd string) (string, []string) {
	tokens := strings.Split(cmd, " ")

	if len(tokens) < 2 {
		return tokens[0], []string{}
	}

	return tokens[0], tokens[1:]

}

func main() {

	builtinCommands["echo"] = 0
	builtinCommands["exit"] = 1
	builtinCommands["type"] = 2

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
