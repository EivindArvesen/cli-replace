package main

import (
	"fmt"
	"log"
    "github.com/BurntSushi/toml"
)

type tomlReplacementConfig struct {
	Find string
	Replace string
}

func readReplacementToml(valuesFileArg string) (string, string) {
	/*
		Read TOML file containing a set of values used to replace values inside text file
	 */
	var config tomlReplacementConfig
	if _, err := toml.DecodeFile(valuesFileArg, &config); err != nil {
		fmt.Println("ERROR: Failed to decode TOML file containing find-and-replace values")
    	log.Fatal(err)
	}

	return config.Find, config.Replace
}

type tomlOutputConfig struct {
	Output string
}

func readOutputToml(configVar string) string {
	/*
		Read TOML configuration file deciding where to store output-file
	 */
	var config tomlOutputConfig
	if _, err := toml.DecodeFile(configVar, &config); err != nil {
		fmt.Println("ERROR: Failed to read configuration file defining output-file")
    	log.Fatal(err)
	}

	return config.Output
}
