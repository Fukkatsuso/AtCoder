package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func dist(x1, y1, x2, y2 int) float64 {
	dx := float64(x1 - x2)
	dy := float64(y1 - y2)
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b, t, v := nextInts(2), nextInts(2), nextInt(), nextInt()
	max := float64(t * v)

	n := nextInt()
	ok := false
	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		if dist(a[0], a[1], x, y)+dist(x, y, b[0], b[1]) <= max {
			ok = true
		}
	}

	if ok {
		puts("YES")
	} else {
		puts("NO")
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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
