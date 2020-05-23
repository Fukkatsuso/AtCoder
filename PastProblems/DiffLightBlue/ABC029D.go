// https://algo-logic.info/digit-dp/#toc_id_3
// 要復習

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

	n := next()
	d := len(n)
	num := make([]int, d)
	for i := range n {
		num[i] = int(n[i] - '0')
	}

	// dp[i][smaller][j]: i桁目まで決め，1がj回現れる数字の個数
	dp := make([][2][11]int, 11)
	dp[0][0][0] = 1

	for i := 0; i < d; i++ {
		for j := 0; j < 10; j++ {
			// smaller: 1->1
			dp[i+1][1][j] += dp[i][1][j] * 9
			dp[i+1][1][j+1] += dp[i][1][j]

			// smaller: 0->1
			if num[i] > 1 {
				dp[i+1][1][j] += dp[i][0][j] * (num[i] - 1)
				dp[i+1][1][j+1] += dp[i][0][j]
			} else if num[i] == 1 {
				dp[i+1][1][j] += dp[i][0][j]
			}

			// smaller: 0->0
			if num[i] == 1 {
				dp[i+1][0][j+1] = dp[i][0][j]
			} else {
				dp[i+1][0][j] = dp[i][0][j]
			}
		}
	}

	ans := 0
	for j := 0; j < 10; j++ {
		ans += (dp[d][0][j] + dp[d][1][j]) * j
	}

	puts(ans)
}
