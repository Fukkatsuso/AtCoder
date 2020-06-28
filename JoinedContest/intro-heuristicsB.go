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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func score(v, d, t int, s [][]int, c, last []int) int {
	last[t] = d + 1
	v += s[d][t]
	for i := 0; i < 26; i++ {
		v -= c[i] * (d + 1 - last[i])
	}
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	D := nextInt()
	c := nextInts(26)
	s := make([][]int, D)
	for i := range s {
		s[i] = nextInts(26)
	}

	last := make([]int, 26)
	v := 0
	for d := 0; d < D; d++ {
		t := nextInt() - 1

		v = score(v, d, t, s, c, last)

		puts(v)
	}
}
