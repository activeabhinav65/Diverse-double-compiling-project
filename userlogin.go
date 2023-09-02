package main

import (
	"fmt"
	"os"
)

func main() {
	cmdLineArguments := os.Args

	userId := cmdLineArguments[1]

	validUserIds := []string{"stark", "hulk", "tony", "hawk"}

	for _, element := range validUserIds {
		if element == userId {
			fmt.Println("Access Granted")
			return
		}
	}

	fmt.Println("Access Denied")
}
