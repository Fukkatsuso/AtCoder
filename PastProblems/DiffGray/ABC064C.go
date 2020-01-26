// 解説AC。問題文読み違えて苦戦
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	n := nextInt()
	col := make([]int, 9)
	for i := 0; i < n; i++ {
		r := nextInt() / 400
		if r >= 9 {
			r = 8
		}
		col[r]++
	}

	maxAns := 0
	for i := 0; i < 8; i++ {
		if col[i] > 0 {
			maxAns++
		}
	}
	maxAns = maxAns + col[8]

	minAns := 0
	for i := 0; i < 8; i++ {
		if col[i] > 0 {
			minAns++
		}
	}
	if minAns == 0 {
		minAns = 1
	}

	puts(minAns, maxAns)

	wt.Flush()
}

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
