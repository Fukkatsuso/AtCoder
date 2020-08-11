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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	x, y, c := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i], c[i] = nextInt(), nextInt(), nextInt()
	}

	low, high := 0.0, 100000000000.0
	for i := 0; i < 100; i++ {
		mid := (low + high) / 2.0
		ok := true
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				can := float64(max(abs(x[j]-x[k]), abs(y[j]-y[k]))) <= mid*(1.0/float64(c[j])+1.0/float64(c[k]))
				ok = ok && can
			}
		}
		if ok {
			high = mid
		} else {
			low = mid
		}
	}

	puts(high)
}
