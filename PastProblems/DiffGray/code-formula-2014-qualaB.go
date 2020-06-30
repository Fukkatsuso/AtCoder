package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b := nextInt(), nextInt()

	pin := make([]int, 10)
	for i := 0; i < a; i++ {
		p := nextInt() - 1
		if p < 0 {
			p = 9
		}
		pin[p] = 1
	}
	for i := 0; i < b; i++ {
		q := nextInt() - 1
		if q < 0 {
			q = 9
		}
		pin[q] = 2
	}

	state := []string{"x", ".", "o"}
	for i := 6; i < 10; i++ {
		putf("%s ", state[pin[i]])
	}
	putf("\n ")
	for i := 3; i < 6; i++ {
		putf("%s ", state[pin[i]])
	}
	putf("\n  ")
	for i := 1; i < 3; i++ {
		putf("%s ", state[pin[i]])
	}
	putf("\n   ")
	for i := 0; i < 1; i++ {
		putf("%s ", state[pin[i]])
	}
	puts()
}
