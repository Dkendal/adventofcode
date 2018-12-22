package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputs []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		txt := scanner.Text()
		val, _ := strconv.Atoi(txt)
		inputs = append(inputs, val)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	sum := 0
	set := make(map[int]bool)
	duplicateFound := false
	firstLoop := true

	for duplicateFound == false {
		for _, val := range inputs {
			sum = sum + val

			if set[sum] == true {
				duplicateFound = true
				fmt.Println("Part 2:", sum)
				break
			}

			set[sum] = true
		}

		if firstLoop {
			fmt.Println("Part 1:", sum)
			firstLoop = false
		}
	}
}
