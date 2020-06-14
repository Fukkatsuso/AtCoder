// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	l := next()
	n := len(l)

	// dp[i][0]: a+b<l, dp[i][1]: a+b=l
	dp := make([][2]int, n+1)
	dp[0][1] = 1
	for i := 1; i <= n; i++ {
		if l[i-1] == '0' {
			dp[i][0] = (3 * dp[i-1][0]) % mod
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][0] = (3*dp[i-1][0] + dp[i-1][1]) % mod
			dp[i][1] = (2 * dp[i-1][1]) % mod
		}
	}

	puts((dp[n][0] + dp[n][1]) % mod)
}
