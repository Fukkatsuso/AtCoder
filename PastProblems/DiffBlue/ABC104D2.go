// 解説AC
// 解き直し

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
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < 4; j++ {
			// ABC文字列に含めない
			if s[i-1] == '?' {
				dp[i][j] += 3 * dp[i-1][j]
			} else {
				dp[i][j] += dp[i-1][j]
			}
			// ABC文字列に含める
			if j > 0 && s[i-1] == []byte("ABC")[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
			if j > 0 && s[i-1] == '?' {
				dp[i][j] += dp[i-1][j-1]
			}
			dp[i][j] %= mod
		}
	}

	puts(dp[n][3])
}
