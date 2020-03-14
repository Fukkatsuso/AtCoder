package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b, c := nextInt(), nextInt(), nextInt()
	ok := true
	if d := c - (a + b); d <= 0 {
		ok = false
	} else {
		ok = 4*a*b < d*d
	}
	if ok {
		puts("Yes")
	} else {
		puts("No")
	}
}

func sqrt(a int) float64 {
	return math.Sqrt(float64(a))
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
