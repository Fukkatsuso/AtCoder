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

	h, w := nextInt(), nextInt()

	// 縦に3分割
	ans := (divFloor(w, 3) - w/3) * h
	// 横に3分割
	ans = min(ans, (divFloor(h, 3)-h/3)*w)
	// 先に横に分割してから縦に分割
	for x := 1; x <= h-1; x++ {
		a, b, c := w*x, (w/2)*(h-x), divFloor(w, 2)*(h-x)
		ans = min(ans, max(a, b, c)-min(a, b, c))
	}
	// 逆バージョン
	for x := 1; x <= w-1; x++ {
		a, b, c := h*x, (h/2)*(w-x), divFloor(h, 2)*(w-x)
		ans = min(ans, max(a, b, c)-min(a, b, c))
	}

	puts(ans)
}
