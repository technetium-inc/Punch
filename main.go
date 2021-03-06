package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// parse the code
func parsePunchCode(code string){
}

// check if the file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// get the filename and read the filecontent
// if the file exists, else
// return the filename along with false
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

// call the readFileContent() function 
// if the length of the os arguments
// is greater than 0
func parseArgumentsExecution(args []string) string {
	// the lenght of the os arguments
	var argumentLength int = len(args)
	if argumentLength != 0 {

		// this is the file content , along with the 
		// result of readFileContent(true/false)
		content, readText = readFileContent(args[0])
		if readText == false {
			os.Exit(3)
		} else {
			parsePunchCode(content)
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
