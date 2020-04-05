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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lowerBound(a []int, x int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if gcd(a[mid], x) == 1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, q := nextInt(), nextInt()
	a, s := nextInts(n), nextInts(q)

	gs := make([]int, n)
	gs[0] = a[0]
	for i := 1; i < n; i++ {
		gs[i] = gcd(gs[i-1], a[i])
	}

	for i := range s {
		lb := lowerBound(gs, s[i])
		if lb == n {
			puts(gcd(gs[n-1], s[i]))
		} else {
			puts(lb + 1)
		}
	}
}
