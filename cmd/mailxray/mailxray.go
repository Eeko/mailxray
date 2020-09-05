package main

import (
	"fmt"
	"os"
	"path/filepath"
	"../../internal/email"
)

// ProcessFilepath is a function to read a string from args and convert it to 
func ProcessFilepath(args []string) []string {
	var files []string
	for _, arg := range args {
		paths, err := filepath.Glob(arg)
		if err != nil {
			panic(err)
		}
		files = append(files, paths...)
	}
	return files
}

// ProcessFile is a function to send each file to the Email analyzer library for processing. 
// Takes in a path which contains a path to a email message file to open.
func ProcessFile(filepath string) {
	filedata, _ := os.Open(filepath)
	message := email.ProcessEmail(filedata)
	fmt.Println(message)
}

// main function
func main() {
	args := os.Args[1:]
	files := ProcessFilepath(args) // read input files to a slice
	for _, filepath := range files {
		ProcessFile(filepath)
	}
}