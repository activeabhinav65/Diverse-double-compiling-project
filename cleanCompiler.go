package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	input := os.Args

	binaryFile := input[3]
	sourceFile := input[4]

	bytes, err := ioutil.ReadFile(sourceFile)

	if err != nil {
		log.Fatal(err)
		return
	}

	sourceCode := string(bytes)

	tempFile := os.TempDir() + "/temp.go"

	err = ioutil.WriteFile(tempFile, []byte(sourceCode), 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile)

	fmt.Print(sourceCode)

	output, err := exec.Command("go", "build", "-o", binaryFile, tempFile).CombinedOutput()

	fmt.Print(string(output))

	if err != nil {
		log.Fatal(err)
	}
}
