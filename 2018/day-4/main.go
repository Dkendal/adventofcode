package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

const (
	timestampFormat = "2006-01-02 15:04"
	inputPattern    = `\[(.+)\] (wakes up|falls asleep|Guard #(\d+) (begins shift))`
)

var pattern = regexp.MustCompile(inputPattern)

type Action int

const (
	start Action = iota
	sleep
	wake
)

type Interval struct {
	t1, t2 time.Time
}

type Event struct {
	ts     time.Time
	action Action
	dat    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Solution(r io.Reader) string {
	s := bufio.NewScanner(r)

	var events []Event

	for s.Scan() {
		t := s.Text()
		matches := pattern.FindStringSubmatch(t)
		ts, e := time.Parse(timestampFormat, matches[1])
		check(e)

		var action Action
		var dat string

		switch matches[2] {
		case "falls asleep":
			action = sleep

		case "wakes up":
			action = wake

		default:
			if matches[4] != "begins shift" {
				panic(0)
			}
			dat = matches[3]
			action = start
		}

		events = append(events, Event{ts, action, dat})
	}

	sort.Slice(events, func(i, j int) bool {
		t1 := events[i].ts
		t2 := events[j].ts
		return t1.Before(t2)
	})

	var id string
	var t1, t2 int
	counters := make(map[string]map[int]int)

	for _, e := range events {
		switch e.action {
		case start:
			id = e.dat

		case sleep:
			t1 = e.ts.Minute()

		case wake:
			t2 = e.ts.Minute()

			for i := t1; i < t2; i++ {
				if counters[id] == nil {
					counters[id] = make(map[int]int)
				}

				counters[id][i]++
			}
		}
	}

	var maxID string
	var maxSum int

	for id, counter := range counters {
		sumSum := 0

		for _, sum := range counter {
			sumSum += sum
		}

		if sumSum > maxSum {
			maxID = id
			maxSum = sumSum
		}
	}

	var maxMinute int
	var maxCount int
	for minute, count := range counters[maxID] {

		if maxCount > count {
			continue
		}

		maxMinute = minute
		maxCount = count
	}

	i, _ := strconv.Atoi(maxID)

	return fmt.Sprintf("%d * %d = %d", i, maxMinute, i*maxMinute)
}

func main() {
	f, e := os.Open("input.txt")

	check(e)

	sol := Solution(f)

	fmt.Println(sol)
}
