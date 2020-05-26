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

	n, a := nextInt2()
	x := nextInts(n)

	// dp[i][j][k]: 1~i番目の範囲で選択する，j枚のカードの合計がkとなる場合の数
	dp := make([][][50*50 + 1]int, n+1)
	for i := range dp {
		dp[i] = make([][50*50 + 1]int, n+1)
	}
	dp[0][0][0] = 1
	maxsum := n * a
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= maxsum; k++ {
				if x[i-1] > k {
					dp[i][j][k] = dp[i-1][j][k]
				} else if j > 0 {
					dp[i][j][k] = dp[i-1][j-1][k-x[i-1]] + dp[i-1][j][k]
				}
			}
		}
	}

	ans := 0
	for j := 1; j <= n; j++ {
		ans += dp[n][j][j*a]
	}

	puts(ans)
}
