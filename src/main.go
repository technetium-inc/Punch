package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// parse the code
func parsePunchCode(code string) {
	statements := strings.Split(code, ";")
	for _, statement := range statements {
		tokens := strings.Split(statement, " ")
        if tokens[0] == "Uppercut[str]" {
			fmt.Println(tokens[1])
		}
    }
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
		b, err := ioutil.ReadFile(string(fileName))
    	// can file be opened?
    	if err != nil {
    	  fmt.Println(err)
    	}
	
    	// convert bytes to string
    	str := string(b)
	
    	// return file data
    	return str, true
	}
	return fileName, false
}

// call the readFileContent() function 
// if the length of the os arguments
// is greater than 0
func parseArgumentsExecution(args []string) string {
	// the length of the os arguments
	var argumentLength int = len(args)
	if argumentLength != 0 {

		// this is the file content , along with the 
		// result of readFileContent(true/false)
		content, readText := readFileContent(args[0])
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
		os.Exit(3)
	} else {
		os.Exit(0)
	}
}
