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

func putSlice(a []int) {
	for i := range a {
		if i == len(a)-1 {
			puts(a[i])
		} else {
			putf("%d ", a[i])
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	p := make([]int, n)
	for i := range p {
		p[i] = nextInt() - 1
	}
	r := make([]int, n)
	for i := range r {
		r[p[i]] = i
	}

	a := make([]int, n)
	for i := range a {
		a[i] = 30000 * (i + 1)
	}
	b := make([]int, n)
	for i := range b {
		b[i] = 30000*(n-i) + r[i]
	}

	putSlice(a)
	putSlice(b)
}
