package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var claimPattern = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func FileClaims(file string) []Claim {
	r, err := os.Open(file)
	defer r.Close()
	check(err)
	return ReadClaims(r)
}

func ReadClaims(r io.Reader) []Claim {
	var claims []Claim
	scan := bufio.NewScanner(r)

	for scan.Scan() {
		line := scan.Text()
		claims = append(claims, parseClaim(line))
	}

	return claims
}

func StringClaims(dat string) []Claim {
	r := strings.NewReader(dat)
	return ReadClaims(r)
}

func OverlappingArea(claims *[]Claim) int {
	sum := 0
	set := make(map[P]int)

	for _, claim := range *claims {
		xMin := claim.pos.x
		yMin := claim.pos.y
		xMax := claim.pos.x + claim.dim.x
		yMax := claim.pos.y + claim.dim.y

		for x := xMin; x < xMax; x++ {
			for y := yMin; y < yMax; y++ {
				p := P{x, y}
				v := set[p]
				set[p]++

				if v == 1 {
					sum++
				}
			}
		}
	}

	return sum
}

func ValidClaims(claims *[]Claim) map[int]bool {
	var (
		valids        = make([]bool, len(*claims))
		visitedPoints = make(map[P][]int)
	)

	for idx, claim := range *claims {
		xMin := claim.pos.x
		yMin := claim.pos.y
		xMax := claim.pos.x + claim.dim.x
		yMax := claim.pos.y + claim.dim.y

		valids[idx] = true

		for x := xMin; x < xMax; x++ {
			for y := yMin; y < yMax; y++ {
				p := P{x, y}
				visitations, ok := visitedPoints[p]
				visitations = append(visitations, idx)
				visitedPoints[p] = visitations

				if ok {
					valids[visitations[0]] = false
				}

				if len(visitations) > 1 {
					valids[idx] = false
				}
			}
		}
	}

	output := make(map[int]bool)

	for idx, ok := range valids {
		if ok {
			n := (*claims)[idx].number
			output[n] = true
		}
	}

	return output
}

func main() {
	claims := FileClaims("input.txt")
	sum := OverlappingArea(&claims)
	fmt.Printf("Part 1: %d\n", sum)

	valids := ValidClaims(&claims)
	fmt.Printf("Part 2: %v\n", valids)
}
