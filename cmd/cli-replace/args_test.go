package main

import "testing"

func TestContainsEmpty(t *testing.T) {
    expected := false
    emptyTest := containsEmpty("Marco", "Polo")
    if emptyTest != expected {
        t.Errorf("ContainsEmpty was incorrect, got: '%t', want: '%t'", emptyTest, expected)
    }
}

func TestDoesNotContainsEmpty(t *testing.T) {
    expected := true
    emptyTest := containsEmpty("Marco", "", "Polo")
    if emptyTest != expected {
        t.Errorf("ContainsEmpty was incorrect, got: '%t', want: '%t'", emptyTest, expected)
    }
}
