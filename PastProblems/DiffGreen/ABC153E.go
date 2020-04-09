// 解説AC
// 他人の提出も参考にした

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

func divFloor(a, b int) int {
	return (a + (b - 1)) / b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	cost := make([]int, 10001)
	for i := 1; i < len(cost); i++ {
		cost[i] = 1000000000
	}

	h, n := nextInt(), nextInt()

	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		for j := 1; j <= h; j++ {
			// HPがjのときにダメージaを与えたときの残りHP
			rest := max(0, j-a)
			cost[j] = min(cost[j], cost[rest]+b)
		}
	}

	puts(cost[h])
}
