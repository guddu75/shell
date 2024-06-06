package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Pwd() (string, error) {

	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return dir, nil

}

func Cd(path string) error {

	homeDir := os.Getenv("HOME")

	if path == "~" {
		return os.Chdir(homeDir)
	}

	curDir, err := Pwd()

	if err != nil {
		fmt.Println(err.Error())
	}

	temp := strings.Split(curDir, "/")

	if strings.HasPrefix(path, "/") {
		return os.Chdir(path)
	}

	steps := strings.Split(path, "/")

	for _, step := range steps {
		if step == "." {
			continue
		} else if step == ".." {
			if len(temp) == 0 {
				return errors.New("no such file present")
			} else {
				temp = temp[:len(temp)-1]
			}
		} else {
			temp = append(temp, step)
		}
	}

	path = strings.Join(temp, "/")

	return os.Chdir(path)

}
