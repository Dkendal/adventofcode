package main

import (
	"testing"
)

func TestOppisite(t *testing.T) {
	assertEqual(t, Oppisite('a', 'A'), true)
	assertEqual(t, Oppisite('A', 'a'), true)
	assertEqual(t, Oppisite('z', 'Z'), true)
	assertEqual(t, Oppisite('Z', 'z'), true)
	assertEqual(t, Oppisite('a', 'a'), false)
	assertEqual(t, Oppisite('A', 'A'), false)
	assertEqual(t, Oppisite('z', 'z'), false)
	assertEqual(t, Oppisite('Z', 'Z'), false)
}

func TestPart1(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"

	actual := Part1(input)

	expected := 10

	assertEqual(t, expected, actual)
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf(`Expected "%v", got: "%v"`, expected, actual)
	}
}
