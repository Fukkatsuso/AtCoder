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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)

	l, r := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		l[i] = l[i-1] ^ a[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1] ^ a[i+1]
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			b[i] = r[i]
		} else if i == n-1 {
			b[i] = l[i]
		} else {
			b[i] = l[i] ^ r[i]
		}
	}

	for i := 0; i < n-1; i++ {
		putf("%d ", b[i])
	}
	puts(b[n-1])
}
