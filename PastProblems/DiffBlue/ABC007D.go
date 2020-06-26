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

// [1,s]に4または9を含む数字がいくつあるか
func digitDP(s string) int {
	n := len(s)
	// [桁][未満/等しい][すでに4,9が出ていない/出ている]
	dp := make([][2][2]int, n+1)
	dp[0][1][0] = 1
	for i := 0; i < n; i++ {
		si := int(s[i] - '0')

		dp[i+1][0][0] += dp[i][0][0] * 8
		dp[i+1][0][1] += dp[i][0][0]*2 + dp[i][0][1]*10

		if si != 4 && si != 9 {
			dp[i+1][1][0] += dp[i][1][0]
		} else {
			dp[i+1][1][1] += dp[i][1][0]
		}
		dp[i+1][1][1] += dp[i][1][1]

		if si >= 1 {
			dp[i+1][0][1] += dp[i][1][1] * si
			if si <= 4 {
				dp[i+1][0][0] += dp[i][1][0] * si
			} else { //si >= 5
				dp[i+1][0][0] += dp[i][1][0] * (si - 1)
				dp[i+1][0][1] += dp[i][1][0]
			}
		}
	}
	return dp[n][0][1] + dp[n][1][1]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, sb := nextInt()-1, next()
	sa := strconv.Itoa(a)

	ans := digitDP(sb) - digitDP(sa)
	puts(ans)
}
