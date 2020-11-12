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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	to := make([]int, n)
	for i := 0; i < m; i++ {
		// x->yに辺を張る
		x, y := getInt()-1, getInt()-1
		to[x] |= (1 << y)
	}

	dp := make([]int, 1<<n)
	dp[0] = 1
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 && i&to[j] == 0 {
				dp[i|(1<<j)] += dp[i]
			}
		}
	}

	puts(dp[(1<<n)-1])
}
