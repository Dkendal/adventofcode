package main

import (
	"testing"
)

func TestClaimOverlapping(t *testing.T) {
	assertOverlapping := func(x Claim, y Claim) {
		if x.Overlapping(y) == false {
			t.Errorf("expected %v to overlap with %v, but it did not", x, y)
		}
	}

	refuteOverlapping := func(x Claim, y Claim) {
		if x.Overlapping(y) == true {
			t.Errorf("expected %v to not overlap with %v, but it did", x, y)
		}
	}

	t.Run("equality property holds", func(t *testing.T) {
		x := Claim{pos: P{0, 0}, dim: P{2, 2}}

		assertOverlapping(x, x)
	})

	t.Run("transitive property holds", func(t *testing.T) {
		x := Claim{pos: P{0, 0}, dim: P{2, 2}}
		y := Claim{pos: P{1, 1}, dim: P{3, 3}}

		assertOverlapping(x, y)
		assertOverlapping(y, x)
	})

	t.Run("adjancent claims don't overlap", func(t *testing.T) {
		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{0, 0}, dim: P{1, 1}},
		)

		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{1, 0}, dim: P{1, 1}},
		)

		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{2, 0}, dim: P{1, 1}},
		)

		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{2, 1}, dim: P{1, 1}},
		)

		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{2, 2}, dim: P{1, 1}},
		)

		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{1, 2}, dim: P{1, 1}},
		)

		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{0, 2}, dim: P{1, 1}},
		)

		refuteOverlapping(
			Claim{pos: P{1, 1}, dim: P{1, 1}},
			Claim{pos: P{0, 1}, dim: P{1, 1}},
		)
	})
}

func TestIndexClaims(t *t.Testing) {
	claims := []Claim{
		Claim{number: 0, pos: P{x: 0, y: 0}, dim: P{x: 1, y: 1}},
		Claim{number: 0, pos: P{x: 0, y: 1}, dim: P{x: 1, y: 2}},
	}

	rects, index := IndexClaims(claims)

	t.Errorf("%v", result)
}
