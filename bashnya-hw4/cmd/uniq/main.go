package main

import (
	"fmt"
	"os"

	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw4/internal/app"
)

const (
	appName = "uniq"
)

func main() {
	if err := app.Launch(os.Args); err != nil {
		fmt.Printf("%s: %s", appName, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
