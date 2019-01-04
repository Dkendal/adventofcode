package main

import (
	"os"
)

func abs(x int64) int64 {
	y := x >> 63
	return (x ^ y) - y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Oppisite(x, y rune) bool {
	return abs(int64(x-y)) == 32
}

func Part1(input string) int {
	s := []rune(input)

	for i := 1; i < len(s); i = max(i+1, 1) {
		j := i - 1

		if Oppisite(s[i], s[j]) == false {
			continue
		}

		s = append(s[:j], s[i+1:]...)
		i -= 2
	}

	return len(s)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, e := os.Open("input.txt")
	check(e)
	defer f.Close()
	s, e := f.Stat()
	check(e)
	b := make([]byte, s.Size())
	_, e = f.Read(b)
	check(e)
	input := string(b)[:len(b) - 1]
	println(Part1(input))
}
