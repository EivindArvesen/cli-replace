package main

import (
	"fmt"
)

var inputFileArg string
var valuesFileArg string
var configFileArg string

var findVal string
var replaceVal string

var outputFile string

func main() {
	fmt.Println("*** cli-replace ***")
	fmt.Println("Parsing args...")

	// Parse command line arguments
	inputFileArg, valuesFileArg, configFileArg = parseArgs(inputFileArg, valuesFileArg, configFileArg)
	if containsEmpty(inputFileArg, valuesFileArg, configFileArg) {
		fmt.Println("Required flags not set; Use the '--help'-flag for more information.")
		return
	}

	fmt.Println("Args parsed!")
	fmt.Println("Reading what to find and replace...")

	// Read what to find and what to replace
	findVal, replaceVal = readReplacementToml(valuesFileArg)

	fmt.Println("Find and replace read!")
	fmt.Println("Reading what file to output to...")

	// Read what file to put the output in
	outputFile = readOutputToml(configFileArg)

	fmt.Println("Output file read!")
	fmt.Println("Performing string replacement on input...")

	// Perform buffered reading from file input file, and writing to output file
	findAndReplace(readFileParameters{inputFile: inputFileArg, callback: writeToFileCallback, outputFile: outputFile})

	fmt.Println("Replacement performed and output!")
}
