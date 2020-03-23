package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := range s {
		s[i] = next()
	}

	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
	}
	if s[0][0] == '#' {
		dp[0][0] = 1
	}
	for i := 1; i < w; i++ {
		dp[0][i] = dp[0][i-1]
		if s[0][i] != s[0][i-1] && s[0][i] == '#' {
			dp[0][i]++
		}
	}
	for i := 1; i < h; i++ {
		dp[i][0] = dp[i-1][0]
		if s[i][0] != s[i-1][0] && s[i][0] == '#' {
			dp[i][0]++
		}
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			down, right := dp[i-1][j], dp[i][j-1]
			if s[i][j] != s[i-1][j] && s[i][j] == '#' {
				down++
			}
			if s[i][j] != s[i][j-1] && s[i][j] == '#' {
				right++
			}
			dp[i][j] = min(down, right)
		}
	}

	puts(dp[h-1][w-1])
}

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
