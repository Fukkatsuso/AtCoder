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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, d := nextInt(), nextInt(), nextInt()
	a := nextInts(m)

	// to[i]: D=1のあみだくじでi番目を選んだときの行き先
	to := make([]int, n)
	for i := range to {
		to[i] = i
	}
	for i := m - 1; i >= 0; i-- {
		to[a[i]-1], to[a[i]] = to[a[i]], to[a[i]-1]
	}

	// dp[i][j]: D=2^iのあみだくじでj番目を選んだときの行き先
	dp := make([][]int, 30)
	dp[0] = make([]int, n)
	for j := range dp[0] {
		dp[0][j] = to[j]
	}
	for i := 1; i < 30; i++ {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	for i := 0; i < n; i++ {
		ans := i
		for j := 0; j < 30; j++ {
			if d&(1<<j) != 0 {
				ans = dp[j][ans]
			}
		}
		puts(ans + 1)
	}
}
