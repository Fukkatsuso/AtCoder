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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, K := next(), nextInt()
	n := len(N)

	// dp[i][j][k]: 上からi桁目, nと同じ(1)/未満(0)の数字と確定, 0でない数字がちょうどk個であるものの個数
	dp := make([][2][4]int, n+2)
	dp[1][0][1] = int(N[0]-'0') - 1
	dp[1][0][0] = 1
	dp[1][1][1] = 1
	for i := 1; i < n; i++ {
		ni := int(N[i] - '0')
		if ni == 0 {
			// 0を使う場合
			for k := 0; k <= 3; k++ {
				dp[i+1][0][k] = dp[i][0][k]
				dp[i+1][1][k] = dp[i][1][k]
			}
			// 0以外を使う場合
			for k := 0; k < 3; k++ {
				dp[i+1][0][k+1] += 9 * dp[i][0][k]
			}
		} else {
			// 0を使う場合
			for k := 0; k <= 3; k++ {
				dp[i+1][0][k] = dp[i][0][k] + dp[i][1][k]
			}
			// 0以外を使う場合
			for k := 0; k < 3; k++ {
				dp[i+1][0][k+1] += 9*dp[i][0][k] + (ni-1)*dp[i][1][k]
				dp[i+1][1][k+1] += dp[i][1][k]
			}
		}
	}

	ans := dp[n][0][K] + dp[n][1][K]
	puts(ans)
}
