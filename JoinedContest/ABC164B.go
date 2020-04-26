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

	a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()

	for t := true; ; t = !t {
		if t {
			c -= b
			if c <= 0 {
				puts("Yes")
				break
			}
		} else {
			a -= d
			if a <= 0 {
				puts("No")
				break
			}
		}
	}
}
