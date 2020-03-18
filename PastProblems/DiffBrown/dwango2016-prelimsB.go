package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	k := nextInts(n - 1)

	l := make([]int, n)
	for i := 1; i < n-1; i++ {
		l[i] = min(k[i-1], k[i])
	}
	if l[1] != k[0] {
		l[0] = k[0]
	} else {
		l[0] = 1
	}
	if l[n-2] != k[n-2] {
		l[n-1] = k[n-2]
	} else {
		l[n-1] = 1
	}

	for i := range l {
		if i < n-1 {
			putf("%d ", l[i])
		} else {
			putf("%d\n", l[i])
		}
	}
}

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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}
