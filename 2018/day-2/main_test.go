package main

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	actual := HammingDistance("abcdefg", "abczefg")

	if actual != 1 {
		t.Errorf("expected: [1], actual: [%v]", actual)
	}
}
