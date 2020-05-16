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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	w := nextInt()
	N, K := nextInt2()
	a, b := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = nextInt2()
	}

	// dp[i][j][k]: k番目まで見る，幅i，j枚
	dp := make([][][]int, w+1)
	for i := range dp {
		dp[i] = make([][]int, N+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, N+1)
		}
	}

	for i := 1; i <= w; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k <= N; k++ {
				if i >= a[k-1] {
					dp[i][j][k] = max(dp[i][j][k-1], dp[i-a[k-1]][j-1][k-1]+b[k-1])
				} else {
					dp[i][j][k] = dp[i][j][k-1]
				}
			}
		}
	}

	puts(dp[w][K][N])
}
