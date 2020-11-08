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

	n := nextInt()
	a := nextInts(n)

	dp := make([]int, n)
	dp[0] = a[0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + a[i]
	}
	rmax := make([]int, n)
	for i := 0; i < n; i++ {
		rmax[i] = dp[i]
		if i > 0 {
			rmax[i] = max(rmax[i], rmax[i-1])
		}
	}

	ans := 0
	x := 0
	for i := 0; i < n; i++ {
		ans = max(ans, x+rmax[i])
		x += dp[i]
	}

	puts(ans)
}
