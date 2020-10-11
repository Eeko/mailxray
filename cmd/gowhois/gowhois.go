package main

import (
	"fmt"
	"os"
	"../../internal/analyzers"

)

// main function
func main() {
	domains := os.Args[1:]
	//domains := ProcessFilepath(args) // read input files to a slice
	for _, domain := range domains {
		lines := analyzers.Whois(domain)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}