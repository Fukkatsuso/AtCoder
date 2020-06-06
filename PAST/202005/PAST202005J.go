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

	n, m := nextInt(), nextInt()
	a := nextInts(m)

	eat := make([]int, n)
	for i := 0; i < m; i++ {
		l, r := 0, n
		for l < r {
			mid := (l + r) / 2
			if a[i] <= eat[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if r == n {
			puts(-1)
		} else {
			eat[r] = a[i]
			puts(r + 1)
		}
	}
}
