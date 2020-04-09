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
	x, y := nextInt(), nextInt()
	a, b := nextInts(n), nextInts(m)

	ans := 0
	for ai, bi, now := 0, 0, 0; ai < n || bi < m; ans++ {
		// aの出発時刻のインデックス
		for ai < n && a[ai] < now {
			ai++
		}
		if ai >= n {
			break
		}
		// aから出発してbに到着
		now = a[ai] + x

		// bからaに向かう
		for bi < m && b[bi] < now {
			bi++
		}
		if bi >= m {
			break
		}
		now = b[bi] + y
	}

	puts(ans)
}
