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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
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

	m, k := getInt(), getInt()

	if k >= (1 << m) {
		puts(-1)
		return
	}

	if m == 0 {
		puts(0, 0)
		return
	} else if m == 1 {
		if k == 0 {
			puts(0, 0, 1, 1)
		} else {
			puts(-1)
		}
		return
	}

	for i := 0; i < (1 << m); i++ {
		if i != k {
			putf("%d ", i)
		}
	}
	putf("%d ", k)
	for i := (1 << m) - 1; i >= 0; i-- {
		if i != k {
			putf("%d ", i)
		}
	}
	puts(k)
}
