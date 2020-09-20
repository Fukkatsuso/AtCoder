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

	n, k := nextInt(), nextInt()
	l, r := make([]int, k), make([]int, k)
	for i := 0; i < k; i++ {
		l[i], r[i] = nextInt(), nextInt()
	}

	// 0-indexed
	dp := make([]int, n)
	// 1-indexed
	a := make([]int, n+1)
	dp[0] = 1
	a[1] = 1
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			L, R := max(0, i-r[j]), max(0, i-l[j]+1)
			dp[i] += (a[R] - a[L] + mod) % mod
			dp[i] %= mod
		}
		a[i+1] = (a[i] + dp[i]) % mod
	}

	puts(dp[n-1])
}
