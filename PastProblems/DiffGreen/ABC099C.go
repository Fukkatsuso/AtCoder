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

	n := nextInt()

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n
	}
	for i := 1; i <= n; i++ {
		for pow := 1; pow <= i; pow *= 6 {
			dp[i] = min(dp[i], dp[i-pow]+1)
		}
		for pow := 1; pow <= i; pow *= 9 {
			dp[i] = min(dp[i], dp[i-pow]+1)
		}
	}

	puts(dp[n])
}
