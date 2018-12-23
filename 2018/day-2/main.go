package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	twos := 0
	threes := 0

	for scanner.Scan() {
		txt := scanner.Text()
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
}
