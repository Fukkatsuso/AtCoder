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

	n, m := nextInt(), nextInt()
	station := make([][]int, n)
	for i := range station {
		station[i] = make([]int, n)
		for j := range station[i] {
			if i == j {
				continue
			}
			station[i][j] = 10000000
		}
	}

	for i := 0; i < m; i++ {
		a, b, t := nextInt()-1, nextInt()-1, nextInt()
		station[a][b], station[b][a] = t, t
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				station[i][j] = min(station[i][j], station[i][k]+station[k][j])
			}
		}
	}

	ans := 10000000
	for i := 0; i < n; i++ {
		// 最悪のケースの最小値が答え
		ans = min(ans, max(station[i]...))
	}

	puts(ans)
}
