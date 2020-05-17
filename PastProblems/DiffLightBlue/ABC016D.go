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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Point struct {
	x, y int
}

func cross(p1, p2, p3, p4 Point) bool {
	ta := (p1.x-p2.x)*(p3.y-p1.y) + (p1.y-p2.y)*(p1.x-p3.x)
	tb := (p1.x-p2.x)*(p4.y-p1.y) + (p1.y-p2.y)*(p1.x-p4.x)
	tc := (p3.x-p4.x)*(p1.y-p3.y) + (p3.y-p4.y)*(p3.x-p1.x)
	td := (p3.x-p4.x)*(p2.y-p3.y) + (p3.y-p4.y)*(p3.x-p2.x)

	return ta*tb < 0 && tc*td < 0
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	ax, ay := nextInt2()
	bx, by := nextInt2()
	pa, pb := Point{ax, ay}, Point{bx, by}
	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}

	c := 0
	for i := 0; i < n; i++ {
		if cross(pa, pb, Point{x[i], y[i]}, Point{x[(i+1)%n], y[(i+1)%n]}) {
			c++
		}
	}

	ans := c/2 + 1
	puts(ans)
}
