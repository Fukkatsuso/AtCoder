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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
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

// copy b to a
func copy(a, b []int) {
	for i := range a {
		a[i] = b[i]
		b[i] = 0
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	a := nextInts(n)

	b := make([]int, n)
	// 高々logN回の操作
	for i := 0; i < min(k, 41); i++ {
		// 累積和で操作結果を求める
		for j := 0; j < n; j++ {
			l, r := max(0, j-a[j]), min(n, j+a[j]+1)
			b[l]++
			if r < n {
				b[r]--
			}
		}
		for j := 1; j < n; j++ {
			b[j] += b[j-1]
		}
		copy(a, b)
	}

	for i := 0; i < n-1; i++ {
		putf("%d ", a[i])
	}
	puts(a[n-1])
}
