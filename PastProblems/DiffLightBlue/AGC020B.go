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

	k := nextInt()
	a := nextInts(k)

	g := func(N int) int {
		for i := range a {
			N -= N % a[i]
		}
		return N
	}

	var low, high int

	low, high = 0, 2+k*max(a...)
	for low < high-1 {
		mid := (low + high) / 2
		if g(mid) >= 2 {
			high = mid
		} else {
			low = mid
		}
	}
	min := high

	low, high = 1, 2+k*max(a...)+1
	for low < high-1 {
		mid := (low + high) / 2
		if g(mid) <= 2 {
			low = mid
		} else {
			high = mid
		}
	}
	max := low

	if min > max {
		puts(-1)
	} else {
		puts(min, max)
	}
}
