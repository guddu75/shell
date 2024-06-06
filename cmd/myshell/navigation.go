package main

import (
	"os"
)

func Pwd() (string, error) {

	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return dir, nil

}

func Cd(path string) error {
	return os.Chdir(path)
}
