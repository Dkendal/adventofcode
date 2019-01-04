package main

import (
	"fmt"
)

type Node struct {
	key  *Interval
	meta meta
	lft  *Node
	rgt  *Node
}

type Interval struct {
	min, max int
}

type meta struct {
	max, sum int
}

func (i *Interval) String() string {
	if i == nil {
		return "[N/A]"
	}
	return fmt.Sprintf("[%d, %d]", i.min, i.max)
}

func (i *Interval) Size() int {
	return i.max - i.min
}

func (i *Interval) Between(x int) bool {
	if i == nil {
		return false
	}
	return i.min <= x && x <= i.max
}

func (t *Node) Query(q int) int {
	if t == nil {
		return 0
	}

	count := 0

	if t.key.Between(q) {
		count++
	}

	if q > t.meta.max {
		return count
	}

	if q > t.key.min {
		count += t.rgt.Query(q)
	}

	count += t.lft.Query(q)

	return count
}

func stringify(t *Node, lvl int) string {
	s := ""
	if t == nil {
		return s
	}
	for i := 0; i < lvl; i++ {
		if i == 0 {
			s += "\n"
		}
		s += "  "
	}
	s += fmt.Sprintf("%v (%+v)", t.key, t.meta)
	s += stringify(t.rgt, lvl+1)
	s += stringify(t.lft, lvl+1)
	return s
}

func (t *Node) Insert(v *Interval) {
	if t.meta.max < v.max {
		(&t.meta).max = v.max
	}

	(&t.meta).sum += v.Size()

	if t.key == nil {
		t.key = v
		return
	}

	if v.min < t.key.min {
		if t.lft == nil {
			t.lft = &Node{}
		}

		t.lft.Insert(v)
		return
	}

	if t.rgt == nil {
		t.rgt = &Node{}
	}

	t.rgt.Insert(v)
}

func (t *Node) String() string {
	return stringify(t, 0)
}
