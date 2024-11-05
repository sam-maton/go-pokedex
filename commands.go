package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println("Help text here...")
	return nil
}

func commandExit() error {
	os.Exit(0)

	return nil
}
