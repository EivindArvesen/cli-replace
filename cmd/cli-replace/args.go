package main

import (
	"flag"
)

func containsEmpty(ss ...string) bool {
	/*
		Check if any of the arguments are empty
	 */
    for _, s := range ss {
        if s == "" {
            return true
        }
    }
    return false
}

func parseArgs(inputFileArg string, valuesFileArg string, configFileArg string) (string, string, string) {
	/*
		Parse the three required command line flags, and return them as strings
	 */
	flag.StringVar(&inputFileArg, "f", "", "help message for input")
	flag.StringVar(&valuesFileArg, "i", "", "help message for values")
	flag.StringVar(&configFileArg, "c", "", "help message for config")

	flag.Parse()

	return inputFileArg, valuesFileArg, configFileArg
}
