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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b := next(), next()
	na, nb := len(a), len(b)

	if na > nb {
		puts("GREATER")
		return
	} else if na < nb {
		puts("LESS")
		return
	}

	for i := range a {
		if a[i] > b[i] {
			puts("GREATER")
			return
		} else if a[i] < b[i] {
			puts("LESS")
			return
		}
	}
	puts("EQUAL")
}
