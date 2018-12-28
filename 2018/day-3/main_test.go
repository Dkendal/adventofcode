package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMain(t *testing.T) {
	c := []Claim{
		Claim{P{0, 0}, P{2, 2}, 0},
		Claim{P{1, 1}, P{2, 2}, 1},
		Claim{P{2, 2}, P{2, 2}, 2},
		Claim{P{4, 4}, P{2, 2}, 3},
	}

	t.Run("OverlappingArea", func(t *testing.T) {
		actual := OverlappingArea(&c)

		expected := 2

		if actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("ValidClaims", func(t *testing.T) {
		actual := ValidClaims(&c)

		expected := map[int]bool{3: true}

		if !cmp.Equal(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
}
