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

	h, w := getInt(), getInt()
	s := make([]string, h)
	for i := range s {
		s[i] = gets()
	}

	dp := make([][]int, h)
	// sum[i][j]: マス(i,j)から{下，右，右下}に移動する場合の(h,w)に到達する組み合わせ(累積和)
	sum := make([][][3]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
		sum[i] = make([][3]int, w)
	}

	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			if i == h-1 && j == w-1 {
				dp[i][j] = 1
				sum[i][j][0], sum[i][j][1], sum[i][j][2] = 1, 1, 1
				continue
			}
			if s[i][j] == '#' {
				continue
			}
			if i < h-1 {
				dp[i][j] += sum[i+1][j][0]
			}
			if j < w-1 {
				dp[i][j] += sum[i][j+1][1]
			}
			if i < h-1 && j < w-1 {
				dp[i][j] += sum[i+1][j+1][2]
			}
			dp[i][j] %= mod

			sum[i][j][0] += dp[i][j]
			sum[i][j][1] += dp[i][j]
			sum[i][j][2] += dp[i][j]
			if i < h-1 {
				sum[i][j][0] += sum[i+1][j][0]
			}
			if j < w-1 {
				sum[i][j][1] += sum[i][j+1][1]
			}
			if i < h-1 && j < w-1 {
				sum[i][j][2] += sum[i+1][j+1][2]
			}
			sum[i][j][0] %= mod
			sum[i][j][1] %= mod
			sum[i][j][2] %= mod
		}
	}

	puts(dp[0][0])
}
