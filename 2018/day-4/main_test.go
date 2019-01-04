package main

import (
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	f, _ := os.Open("example.txt")
	e := NewEvents(f)
	actual1, actual2 := Solution(e)

	expected1 := 240
	expected2 := 4455

	if expected1 != actual1 {
		t.Errorf(`Expected "%v", got: "%v"`, expected1, actual1)
	}

	if expected2 != actual2 {
		t.Errorf(`Expected "%v", got: "%v"`, expected2, actual2)
	}
}
