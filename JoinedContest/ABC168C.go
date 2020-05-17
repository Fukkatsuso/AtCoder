package main

import (
	"bufio"
	"fmt"
	"math"
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

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b, h, m := nextInt4()
	deg := float64(6*m) - (float64(30*h) + float64(m)/2.0)
	if deg < 0 {
		deg *= -1
	}
	if deg > 180.0 {
		deg = 360.0 - deg
	}
	pi := (math.Pi / 180.0) * deg

	dist := math.Sqrt(float64(a*a) + float64(b*b) - 2.0*float64(a)*float64(b)*math.Cos(pi))
	puts(dist)
}
