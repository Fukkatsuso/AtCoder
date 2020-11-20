// 写経AC
// https://algo-logic.info/abc155e/
// わかりやすい解説
// https://www.creativ.xyz/abc155-e/

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

	n := gets()
	l := len(n)

	dp := make([][2]int, l+1)
	dp[0][1] = 1
	for i := 0; i < l; i++ {
		ni := int(n[i] - '0')
		dp[i+1][0] = min(dp[i][0]+ni, dp[i][1]+10-ni)
		dp[i+1][1] = min(dp[i][0]+ni+1, dp[i][1]+10-ni-1)
	}

	puts(dp[l][0])
}
