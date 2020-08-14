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
	inf            = 1000000000
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

	h, w := nextInt(), nextInt()
	a := make([]string, h)
	for i := range a {
		a[i] = next()
	}

	cost := make([][]int, h)
	for i := range cost {
		cost[i] = make([]int, w)
		for j := range cost[i] {
			cost[i][j] = inf
		}
	}

	// タテ方向を見る
	// 上から下
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			if i > 0 {
				cost[i][j] = cost[i-1][j] + 1
			}
			if a[i][j] == '#' {
				cost[i][j] = 0
			}
		}
	}
	// 下から上
	for j := 0; j < w; j++ {
		for i := h - 2; i >= 0; i-- {
			cost[i][j] = min(cost[i][j], cost[i+1][j]+1)
		}
	}

	// ヨコ方向を見る
	// 左から右
	for i := 0; i < h; i++ {
		for j := 1; j < w; j++ {
			cost[i][j] = min(cost[i][j], cost[i][j-1]+1)
		}
	}
	// 右から左
	for i := 0; i < h; i++ {
		for j := w - 2; j >= 0; j-- {
			cost[i][j] = min(cost[i][j], cost[i][j+1]+1)
		}
	}

	ans := 0
	for i := range cost {
		ans = max(ans, max(cost[i]...))
	}
	puts(ans)
}
