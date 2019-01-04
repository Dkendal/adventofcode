package main

import (
	"os"
	"regexp"
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

func React(input string, c chan int) {
	s := []rune(input)

	for i := 1; i < len(s); i = max(i+1, 1) {
		j := i - 1

		if Oppisite(s[i], s[j]) == false {
			continue
		}

		s = append(s[:j], s[i+1:]...)
		i -= 2
	}

	c <- len(s)
}

func ReduceReact(input string, c chan int) {
	m := make(map[rune]chan int)

	for r := 'a'; r <= 'z'; r++ {
		ch := make(chan int)
		m[r] = ch
		pat := "[" + string(r) + string(r-32) + "]"
		reg := regexp.MustCompile(pat)
		in := reg.ReplaceAllString(input, "")
		go React(in, ch)
	}

	min := struct {
		length int
		char   rune
	}{
		-1,
		'a',
	}

	for r, ch := range m {
		length := <-ch

		if length < min.length || min.length == -1 {
			min.length = length
			min.char = r
		}
	}

	println("done", string(min.char))
	c <- min.length
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
	input := string(b)[:len(b)-1]

	c := make(chan int)
	go React(input, c)
	a := <-c
	println(a)

	c = make(chan int)
	go ReduceReact(input, c)
	a = <-c
	println(a)
}
