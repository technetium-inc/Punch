package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

// parse the code
func Punch(code string) {
	lives := 100
	symTab := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
	}
	statements := strings.Split(code, ";\r\n")
	for _, statement := range statements {
		tokens := strings.Split(statement, " ")
		switch tokens[0] {
		case "TrashTalkLn[str]":
			lives -= 5
			if lives > 0 {
				fmt.Println(strings.Replace(strings.Trim(tokens[1], "\""), "\\s", " ", -1))
			}

		case "TrashTalk[str]":
			lives -= 5
			if lives > 0 {
				fmt.Print(strings.Replace(strings.Trim(tokens[1], "\""), "\\s", " ", -1))
			}
		case "TrashTalk[int]":
			lives -= 5
			if lives > 0 {
				fmt.Print(tokens[1])
			}
		case "TrashTalkLn[int]":
			lives -= 5
			if lives > 0 {
				fmt.Println(tokens[1])
			}
		case "FaceSlap[var,int]":
			lives -= 4
			if lives > 0 {
				symTab[tokens[1]], _ = strconv.Atoi(tokens[2])
			}
		case "TrashTalk[var]":
			lives -= 5
			if lives > 0 {
				fmt.Print(symTab[tokens[1]])
			}
		case "TrashTalkLn[var]":
			lives -= 5
			if lives > 0 {
				fmt.Println(symTab[tokens[1]])
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
			Punch(content)
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
