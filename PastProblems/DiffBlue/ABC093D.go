// 解説AC

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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

// ab > c^2 を満たす最大のc
func sqrtFloor(a, b int) int {
	l, r := 0, 1000000001
	for l < r-1 {
		c := (l + r) / 2
		if a*b > c*c {
			l = c
		} else {
			r = c
		}
	}
	return l
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	q := getInt()

	for i := 0; i < q; i++ {
		a, b := getInt(), getInt()
		if a > b {
			a, b = b, a
		}
		if b-a <= 1 {
			puts(2*a - 2)
		} else if c := sqrtFloor(a, b); c*(c+1) >= a*b {
			puts(2*c - 2)
		} else {
			puts(2*c - 1)
		}
	}
}
