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

	n, q := nextInt(), nextInt()
	turn := make([]int, n+1)
	for i := 0; i < q; i++ {
		l, r := nextInt()-1, nextInt()-1
		turn[l]++
		turn[r+1]--
	}

	for i, cnt := 0, 0; i < n; i++ {
		cnt += turn[i]
		putf("%d", cnt%2)
	}
	puts()
}
