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
	mod            = 1000000007
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

	s := gets()
	n := len(s)

	dp := make([][4]int, n+1)
	dp[n][3] = 1
	for i := n - 1; i >= 0; i-- {
		// j == 3
		if s[i] == '?' {
			dp[i][3] = (3 * dp[i+1][3]) % mod
		} else {
			dp[i][3] = dp[i+1][3]
		}
		// j < 3
		for j := 2; j >= 0; j-- {
			m1, m2 := 1, 0
			if s[i] == '?' {
				m1, m2 = 3, 1
			} else if int(s[i]-'A') == j {
				m2 = 1
			}
			dp[i][j] = m1*dp[i+1][j] + m2*dp[i+1][j+1]
			dp[i][j] %= mod
		}
	}

	ans := dp[0][0]
	puts(ans)
}
