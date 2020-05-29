// ansの計算方法をミスっていた
// 他の人の解答を見て修正

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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, W := nextInt(), nextInt()
	w, v := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		w[i], v[i] = nextInt(), nextInt()
	}

	// dp[i][j][k]: i番目まで見て，j個のモノを使って重さk以下を満たすときの価値の最大値
	dp := make([][][301]int, N+1)
	for i := range dp {
		dp[i] = make([][301]int, N+1)
	}

	for i := 1; i <= N; i++ {
		// 0 <= wi <= 3
		wi, vi := w[i-1]-w[0], v[i-1]
		for j := 1; j <= N; j++ {
			for k := 0; k <= 300; k++ {
				if k >= wi {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-1][k-wi]+vi)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	ans := 0
	for j := 1; j <= N; j++ {
		rem := W - w[0]*j
		if rem >= 0 {
			ans = max(ans, dp[N][j][min(rem, 300)])
		}
	}

	puts(ans)
}
