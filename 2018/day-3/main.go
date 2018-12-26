package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var _ = fmt.Printf
var _ = strconv.ParseInt
var claimPattern = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Rect struct {
	xMin, yMin, xMax, yMax int
}

// P is a 2d vector
type P struct {
	x, y int
}

// Claim is the primary construct for the problem
type Claim struct {
	pos    P
	dim    P
	number int
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func (t Claim) String() string {
	return fmt.Sprintf(
		"<#%d @ %d,%d: %dx%d>",
		t.number,
		t.pos.x,
		t.pos.x,
		t.dim.x,
		t.dim.y,
	)
}

func (t Claim) topLeft() P {
	return t.pos
}

func (t Claim) topRight() P {
	return P{
		x: t.pos.x + t.dim.x,
		y: t.pos.y,
	}
}

func (t Claim) bottomLeft() P {
	return P{
		x: t.pos.x,
		y: t.pos.y + t.dim.y,
	}
}

func (t Claim) bottomRight() P {
	return P{
		x: t.pos.x + t.dim.x,
		y: t.pos.y + t.dim.y,
	}
}

func (t Claim) left() int   { return t.pos.x }
func (t Claim) right() int  { return t.pos.x + t.dim.x }
func (t Claim) top() int    { return t.pos.y }
func (t Claim) bottom() int { return t.pos.y + t.dim.y }

func (t Claim) inside(b P) bool {
	return b.x > t.left() &&
		b.x < t.right() &&
		b.y > t.top() &&
		b.y < t.bottom()
}

// Overlapping is:
// if any part of x is inside the boundary of y == if
// any corner of x is within y or any corner of y is within x
func (t Claim) Overlapping(y Claim) bool {
	return (t.pos == y.pos && t.dim == t.dim) ||
		y.inside(t.topLeft()) ||
		y.inside(t.topRight()) ||
		y.inside(t.bottomLeft()) ||
		y.inside(t.bottomRight())
}

func parseClaim(str string) Claim {
	match := claimPattern.FindStringSubmatch(str)

	number := atoi(match[1])
	x := atoi(match[2])
	y := atoi(match[3])
	w := atoi(match[4])
	h := atoi(match[5])

	return Claim{
		number: number,
		pos:    P{x, y},
		dim:    P{w, h},
	}
}

func getClaims(file string) []Claim {
	dat, err := os.Open(file)

	defer dat.Close()

	check(err)

	scan := bufio.NewScanner(dat)

	var claims []Claim

	for scan.Scan() {
		line := scan.Text()
		claims = append(claims, parseClaim(line))
	}

	return claims
}

IndexRectangles(claims []Claims) {
}

func main() {
	claims := getClaims("./input.txt")

	indexRectangles(claims)
}
