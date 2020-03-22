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

	n := nextInt()
	ng := nextInts(3)

	for i := range ng {
		if ng[i] == n {
			puts("NO")
			return
		}
	}

	dp := make([]bool, n+1)
	cnt := make([]int, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		cnt[i] = n + 1
		if i == ng[0] || i == ng[1] || i == ng[2] {
			dp[i] = false
			continue
		}

		if dp[i-1] {
			dp[i] = true
			cnt[i] = cnt[i-1] + 1
		}
		if i >= 2 && dp[i-2] {
			dp[i] = true
			cnt[i] = min(cnt[i], cnt[i-2]+1)
		}
		if i >= 3 && dp[i-3] {
			dp[i] = true
			cnt[i] = min(cnt[i], cnt[i-3]+1)
		}
	}

	if dp[n] && cnt[n] <= 100 {
		puts("YES")
	} else {
		puts("NO")
	}
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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}
