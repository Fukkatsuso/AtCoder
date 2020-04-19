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

	s, k := next(), nextInt()
	n := len(s)

	toa := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] > 'a' {
			toa[i] = int('z'-s[i]) + 1
		}
	}

	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = s[i]
	}
	for i := 0; i < n && k > 0; i++ {
		if k >= toa[i] {
			ans[i] = 'a'
			k -= toa[i]
		}
	}

	if k > 0 {
		ans[n-1] += byte(k % 26)
	}

	puts(string(ans))
}
