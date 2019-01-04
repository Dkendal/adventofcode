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

type Event struct {
	ts     time.Time
	action Action
	dat    int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewEvents(r io.Reader) []Event {
	s := bufio.NewScanner(r)

	var events []Event

	for s.Scan() {
		t := s.Text()
		matches := pattern.FindStringSubmatch(t)
		ts, e := time.Parse(timestampFormat, matches[1])
		check(e)

		if matches == nil {
			panic(0)
		}

		var action Action
		var dat int

		switch matches[2] {
		case "falls asleep":
			action = sleep

		case "wakes up":
			action = wake

		default:
			if matches[4] != "begins shift" {
				panic(0)
			}
			var e error
			dat, e = strconv.Atoi(matches[3])
			check(e)
			action = start
		}

		events = append(events, Event{ts, action, dat})
	}

	sort.Slice(events, func(i, j int) bool {
		t1 := events[i].ts
		t2 := events[j].ts
		return t1.Before(t2)
	})

	return events
}

func EachInterval(events *[]Event, fn func(id int, t1, t2 int)) {
	var id, t1, t2 int

	for _, e := range *events {
		switch e.action {
		case start:
			id = e.dat

		case sleep:
			t1 = e.ts.Minute()

		case wake:
			t2 = e.ts.Minute()

			fn(id, t1, t2-1)

		default:
			panic("unhandled case")
		}
	}
}

func NewForest(events []Event) map[int]*Node {
	var m = make(map[int]*Node)

	EachInterval(&events, func(id, t1, t2 int) {
		node, ok := m[id]

		if !ok {
			node = &Node{}
			m[id] = node
		}

		node.Insert(&Interval{t1, t2})
	})

	return m
}

func Part1(events []Event) int {
	m := NewForest(events)

	var id int
	var maxSum int
	for i, n := range m {
		if maxSum < n.meta.sum {
			maxSum = n.meta.sum
			id = i
		}
	}

	n := m[id]

	var minute, maxCount int
	for i := 0; i < 60; i++ {
		count := n.Query(i)
		if count > maxCount {
			maxCount = count
			minute = i
		}
	}

	return id * minute
}

func Part2(events []Event) int {
	m := NewForest(events)

	var choice struct {
		minute int
		count  int
		id     int
	}

	for id, node := range m {
		for minute := 0; minute < 60; minute++ {
			count := node.Query(minute)
			if count > choice.count {
				choice.count = count
				choice.minute = minute
				choice.id = id
			}
		}
	}

	return choice.id * choice.minute
}

func Solution(events []Event) (int, int) {
	return Part1(events), Part2(events)
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	e := NewEvents(f)
	part1, part2 := Solution(e)
	fmt.Println(part1)
	fmt.Println(part2)
}
