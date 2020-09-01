// 解説AC
// 参照
// https://www.hamayanhamayan.com/entry/2020/06/21/153453

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
	mod            = 998244353
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

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b, c, d := nextInt4()

	dp := make([][]int, c+1)
	for i := range dp {
		dp[i] = make([]int, d+1)
	}

	dp[a][b] = 1
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			dp[i][j] += (dp[i-1][j] * j) % mod
			dp[i][j] += (dp[i][j-1] * i) % mod
			dp[i][j] -= (dp[i-1][j-1] * (i - 1) * (j - 1)) % mod
			dp[i][j] = (dp[i][j] + mod) % mod
		}
	}

	puts(dp[c][d])
}
