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

	n, s := nextInt(), next()

	l, r := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == ')' {
			if l == 0 {
				putf("(")
			} else {
				l--
			}
		} else {
			l++
		}
	}
	putf("%s", s)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			if r == 0 {
				putf(")")
			} else {
				r--
			}
		} else {
			r++
		}
	}
	putf("\n")
}
