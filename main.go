package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func readFileContent(fileName string) (string, bool) {
	if fileExists(fileName) {
		fileContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal("Failed to read the file")
		} else {
			return string(fileContent), true
		}
	}
	return fileName, false
}

func parseArgumentsExecution(args []string) string {
	var argumentLength int = len(args)
	if argumentLength != 0 {

		// this is the file content i guess
		content, readText = readFileContent(args[0])
		if readText == false {
			os.Exit(3)
		}
		return args[0]
	}

	return "nil"
}

func main() {
	cliArguments := os.Args[1:]

	arguments := parseArgumentsExecution(cliArguments)
	if arguments == "nil" {

	} else {
		fmt.Println(arguments)
	}
}
