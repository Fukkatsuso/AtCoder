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
	a := nextInts(n)

	ok := true
	for i := 1; ok && i < n; i++ {
		ok = ok && (a[i-1] <= a[i])
	}
	if ok {
		puts(0)
		return
	}

	plus := true
	mx, mxI := 0, 0
	for i := range a {
		if mx < abs(a[i]) {
			plus = (a[i] > 0)
			mx, mxI = abs(a[i]), i+1
		}
	}

	// 各項1回ずつ符号付きmxを加算
	puts(2*n - 1)
	for i := 1; i <= n; i++ {
		puts(mxI, i)
	}
	// n-1回の操作で累積和計算
	if plus {
		for i := 1; i < n; i++ {
			puts(i, i+1)
		}
	} else {
		for i := n; i > 1; i-- {
			puts(i, i-1)
		}
	}
}
