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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
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

	n, m := getInt(), getInt()
	a, b := getInts(n), getInts(m)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = 1 << 50
		}
	}
	dp[0][0] = 0

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			}
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
			}
			if i > 0 && j > 0 {
				if a[i-1] == b[j-1] {
					dp[i][j] = min(dp[i][j], dp[i-1][j-1])
				} else {
					// 2つとも消すより不一致の1ペナを被る方が得
					dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
				}
			}
		}
	}

	puts(dp[n][m])
}
