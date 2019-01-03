package main

import (
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	f, _ := os.Open("example.txt")

	expected := "10 * 24 = 240"
	actual := Solution(f)

	if expected != actual {
		t.Errorf(`Expected "%s", got: "%s"`, expected, actual)
	}
}
