package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	twos := 0
	threes := 0
	var list []string

	for scanner.Scan() {
		txt := scanner.Text()
		list = append(list, txt)
		runes := make(map[rune]int)
		hasTwo := false
		hasThree := false

		for _, char := range txt {
			runes[char]++
		}

		for _, count := range runes {
			if count == 2 {
				hasTwo = true
			}
			if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	checksum := twos * threes

	println("[Part 1] checksum:", checksum)

	length := len(list)

	for i, x := range list {
		for _, y := range list[i+1 : length] {
			if HammingDistance(x, y) == 1 {
				fmt.Printf("[Part 2] matching code: ")

				for i := range x {
					if x[i] == y[i] {
						fmt.Printf("%c", x[i])
					}
				}

				fmt.Printf("\n")
			}
		}
	}
}

// HammingDistance find the edit distance.
func HammingDistance(x string, y string) int {
	if len(x) != len(y) {
		panic("expected strings to be the same length")
	}

	dist := 0

	for i := range x {
		if x[i] != y[i] {
			dist++
		}
	}

	return dist
}
