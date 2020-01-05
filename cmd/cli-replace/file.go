package main

import (
    "fmt"
    "io"
    "os"
    "log"
    "strings"
)

type callbackFunc func(string, string)

func printCallback(args ...string) {
    /*
        Callback that prints input; for debugging
     */
	fmt.Print(args[0])
}

func WriteToFile(filename string, data string) error {
    /*
        Prints any string of text to a file safely by checking for errors and syncing at the end
     */

    // Create file, if it does not exist
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("ERROR: Failed to create file to write to")
        return err
    }
    defer file.Close()

    // Write string to file
    _, err = io.WriteString(file, data)
    if err != nil {
        fmt.Println("ERROR: Failed to write string to file")
        return err
    }
    return file.Sync()
}

func writeToFileCallback(input string, outputFile string) {
	err := WriteToFile(outputFile, input)
    if err != nil {
        fmt.Println("ERROR: Failed writing output to file via callback")
        log.Fatal(err)
    }
}

type readFileParameters struct {
	inputFile string
	callback callbackFunc
    outputFile string
}

func findAndReplace(params readFileParameters) bool {
    /*
        Perform find and replace.

        We do buffered reading of file (to avoid memory issues on large files),
        and convert buffer to string using string builder in order to avoid excess copying

        TODO: This should maybe be broken up into smaller pieces...
     */

    // Check that input file can be opened
	file, err := os.Open(params.inputFile)
    if err != nil {
    	// Can't open file
        log.Fatal(err)
    }
    defer file.Close()

    buf := make([]byte, 32*1024) // define your buffer size here.
    var stringBuilder strings.Builder

    // Do buffered reading of file (to avoid memory issues on large files)
    for {
        n, err := file.Read(buf)

        if n > 0 {
            fmt.Fprintf(&stringBuilder, "%s", string(buf[:n])) // your read buffer.

            // Convert buffer content to string and replace values
            bufferString := stringBuilder.String()
            processedString := strings.ReplaceAll(bufferString, findVal, replaceVal)

            // If it is set, call optional callback; The callback would deal with output (e.g. print, write to file)
            if params.callback!= nil {
            	params.callback(processedString, params.outputFile)
            }
        }

        // End Of File - Done reading.
        if err == io.EOF {
            break
        }
        // Error-handling
        if err != nil {
            fmt.Println("ERROR: Failed buffered reading of input file")
            log.Printf("read %d bytes: %v", n, err)
            break
        }
    }
    return true
}
