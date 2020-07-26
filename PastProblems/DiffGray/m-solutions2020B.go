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

	a, b, c := nextInt(), nextInt(), nextInt()
	k := nextInt()

	for i := 0; i < k; i++ {
		if a >= b {
			b *= 2
		} else if b >= c {
			c *= 2
		}
	}

	if a < b && b < c {
		puts("Yes")
	} else {
		puts("No")
	}
}
