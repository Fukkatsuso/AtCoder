// 0点

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
	for i := 1; i <= D; i++ {
		t := 0
		add := -(1 << 60)
		for j := 0; j < 26; j++ {
			old := last[j]
			last[j] = i
			diff := s[i-1][j]
			for k := range c {
				diff += c[k] * (i - last[k])
			}
			if diff > add {
				t = j
				add = diff
			}
			last[j] = old
		}
		puts(t + 1)
		last[t] = i
	}
}
