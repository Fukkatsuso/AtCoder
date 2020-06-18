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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1 << 40

	n, m := nextInt(), nextInt()

	a, d := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = nextInt()
		b := nextInt()
		for j := 0; j < b; j++ {
			c := nextInt() - 1
			d[i] |= (1 << uint(c))
		}
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, 1<<uint(n))
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for i := 0; i < m; i++ {
		for j := 0; j < (1 << uint(n)); j++ {
			dp[i+1][j] = min(dp[i+1][j], dp[i][j])
			dp[i+1][j|d[i]] = min(dp[i+1][j|d[i]], dp[i][j]+a[i])
		}
	}

	if dp[m][(1<<uint(n))-1] == inf {
		puts(-1)
	} else {
		puts(dp[m][(1<<uint(n))-1])
	}
}
