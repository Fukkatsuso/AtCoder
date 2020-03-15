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

	h, w := nextInt(), nextInt()
	a := make([]string, h)
	for i := 0; i < h; i++ {
		a[i] = next()
	}

	ok := true
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if a[i][j-1] == '#' && a[i-1][j] == '#' {
				ok = false
			}
		}
	}

	if ok {
		puts("Possible")
	} else {
		puts("Impossible")
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
