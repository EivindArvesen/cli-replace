package main

import "testing"

func TestFindAndReplace(t *testing.T) {
	expected := true
	file_test := findAndReplace(readFileParameters{inputFile: "../../test/data/myInput.txt", callback: writeToFileCallback, outputFile: "ouptut.txt"})
    if file_test != expected {
        t.Errorf("TestFindAndReplace was incorrect, got: '%t', want: '%t'", file_test, expected)
    }
}
