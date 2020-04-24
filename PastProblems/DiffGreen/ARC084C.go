// 解説の解法とは違うが一応AC

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

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = nextFloat64()
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

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const MAX = 1000000001

	n := nextInt()
	a, b, c := nextInts(n), nextInts(n), nextInts(n)
	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))
	sort.Sort(sort.IntSlice(c))

	sb, sc := make([]int, n), make([]int, n)
	for ai, bi := 0, 0; bi < n; bi++ {
		for ai < n && a[ai] < b[bi] {
			sb[bi]++
			ai++
		}
		if bi > 0 {
			sb[bi] += sb[bi-1]
		}
	}
	for bi, ci := 0, 0; ci < n; ci++ {
		for bi < n && b[bi] < c[ci] {
			sc[ci] += sb[bi]
			bi++
		}
		if ci > 0 {
			sc[ci] += sc[ci-1]
		}
	}

	puts(sum(sc))
}
