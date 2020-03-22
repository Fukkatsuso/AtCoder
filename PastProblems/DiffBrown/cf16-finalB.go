package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	max := n
	for l, r := 1, n; l <= r; {
		mid := (r + l) / 2
		max = mid
		if s := sum(mid); l == r || s == n {
			break
		} else if s < n {
			l = mid + 1
		} else {
			r = mid
		}
	}

	rm := sum(max) - n
	for i := 1; i <= max; i++ {
		if i != rm {
			puts(i)
		}
	}
}

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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func sum(a int) int {
	return (a * (a + 1)) / 2
}
