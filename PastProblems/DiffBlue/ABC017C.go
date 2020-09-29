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

	n, m := nextInt(), nextInt()
	dp := make([]int, m)
	sum := 0
	for i := 0; i < n; i++ {
		l, r, s := nextInt()-1, nextInt()-1, nextInt()
		dp[l] += s
		if r+1 < m {
			dp[r+1] -= s
		}
		sum += s
	}
	for i := 1; i < m; i++ {
		dp[i] += dp[i-1]
	}

	ans := sum - min(dp...)
	puts(ans)
}
