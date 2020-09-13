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
	mod            = 1000000007
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	dp := make([][4]int, n+1)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] * 8
		dp[i][1] = dp[i-1][0] + dp[i-1][1]*9
		dp[i][2] = dp[i-1][0] + dp[i-1][2]*9
		dp[i][3] = dp[i-1][1] + dp[i-1][2] + dp[i-1][3]*10
		for j := range dp[i] {
			dp[i][j] %= mod
		}
	}

	puts(dp[n][3])
}
