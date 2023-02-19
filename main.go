package main

import (
	"flag"     // package for command-line flag parsing
	"fmt"      // package for console output
	"os"       // package for interacting with the operating system
	"strconv"  // package for string to integer conversion
	"strings"  // package for string manipulation
)

func main() {
	// define command-line flags
	inputFile := flag.String("i", "", "input file")
	outputFile := flag.String("o", "", "output file")
	flag.Parse()

	var decimalNumbers []string

	if *inputFile != "" {
		// read decimal numbers from input file
		fileContent, err := os.ReadFile(*inputFile)
		if err != nil {
			// print error message to stderr and exit with non-zero status code
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}
		// split the file content into whitespace-separated fields (i.e., decimal numbers)
		decimalNumbers = strings.Fields(string(fileContent))
	} else {
		// read decimal numbers from command-line arguments
		decimalNumbers = flag.Args()
	}

	// convert decimal numbers to ASCII text
	var asciiText string
	for _, decimalStr := range decimalNumbers {
		// convert the decimal number string to an integer
		decimal, err := strconv.Atoi(decimalStr)
		if err != nil {
			// print error message to stderr and skip the invalid number
			fmt.Fprintf(os.Stderr, "invalid decimal number: %s\n", decimalStr)
			continue
		}
		// convert the decimal number to its ASCII text equivalent and append it to the output string
		asciiText += string(decimal)
	}

	if *outputFile != "" {
		// write the output string to the specified output file
		err := os.WriteFile(*outputFile, []byte(asciiText), 0644)
		if err != nil {
			// print error message to stderr and exit with non-zero status code
			fmt.Fprintf(os.Stderr, "error writing file: %v\n", err)
			os.Exit(1)
		}
	} else {
		// print the output string to the console
		fmt.Println(asciiText)
	}
}