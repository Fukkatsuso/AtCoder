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

func divFloor(a, b int) int {
	return (a + (b - 1)) / b
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

	n, a, b := nextInt(), nextInt(), nextInt()
	h := nextInts(n)

	// x回の攻撃で倒せるか
	canKillAll := func(x int) bool {
		acnt := 0
		for i := 0; i < n; i++ {
			acnt += divFloor(max(0, h[i]-b*x), a-b)
		}
		return acnt <= x
	}

	low, high := 0, 1000000000
	for low+1 < high {
		mid := (low + high) / 2
		if canKillAll(mid) {
			high = mid
		} else {
			low = mid
		}
	}

	puts(high)
}
