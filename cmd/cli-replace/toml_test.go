package main

import "testing"

func TestReadReplacementToml(t *testing.T) {
    expected_find := "boat"
    expected_replace := "car"
    find_test, replace_test := readReplacementToml("../../test/data/myValues.toml")
    if find_test != expected_find {
        t.Errorf("ReadReplacementToml was incorrect, got: '%s', want: '%s'", find_test, expected_find)
    }
    if replace_test != expected_replace {
        t.Errorf("ReadReplacementToml was incorrect, got: '%s', want: '%s'", replace_test, expected_replace)
    }
}


func TestReadOutputToml(t *testing.T) {
    expected_output := "output.txt"
    output_test := readOutputToml("../../test/data/myApp.toml")
    if output_test != expected_output {
        t.Errorf("ReadReplacementToml was incorrect, got: '%s', want: '%s'", output_test, expected_output)
    }
}
