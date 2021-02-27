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

	n := getInt()
	s := make([]string, n)
	for i := range s {
		s[i] = gets()
	}

	dp := make([][2]int, n+1)
	dp[0][0], dp[0][1] = 1, 1
	for i := 1; i <= n; i++ {
		if s[i-1] == "AND" {
			dp[i][0] = 2*dp[i-1][0] + dp[i-1][1]
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][0] + 2*dp[i-1][1]
		}
	}

	puts(dp[n][1])
}
