package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n, k := nextInt(), nextInt()
	a, f := nextInts(n), nextInts(n)
	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.Reverse(sort.IntSlice(f)))

	cost := func(t int) int {
		res := 0
		for i := 0; i < n; i++ {
			res += max(0, a[i]-t/f[i])
		}
		return res
	}

	low, high := -1, 1000000000000
	for low+1 < high {
		mid := (low + high) / 2
		if cost(mid) <= k {
			high = mid
		} else {
			low = mid
		}
	}

	puts(high)
}
