package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	wt = bufio.NewWriter(os.Stdout)
)

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	defer wt.Flush()

	n := 20
	a := make([]int, n)
	a[0], a[1], a[2] = 100, 100, 200
	for i := 3; i < n; i++ {
		a[i] = a[i-1] + a[i-2] + a[i-3]
	}

	puts(a[n-1])
}
