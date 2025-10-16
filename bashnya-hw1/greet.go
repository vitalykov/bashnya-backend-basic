package main

import (
	"fmt"
	"os/user"
	"runtime"
)

func findUserName() (string, error) {
	user, err := user.Current()
	var name string
	if err != nil {
		return name, err
	}
	name = user.Name
	return name, nil
}

func Greet() {
	username, err := findUserName()
	if err != nil {
		username = "some entity"
	}
	fmt.Printf("Hello, %s!\n", username)
	fmt.Printf("Running from %s\n", runtime.GOOS)
}
