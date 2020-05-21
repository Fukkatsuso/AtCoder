// 解説AC

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

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b, c := nextFloat64(), nextFloat64(), nextFloat64()

	f := func(t float64) float64 {
		return a*t + b*math.Sin(c*t*math.Pi)
	}

	low, high := 0.0, 200.0
	for i := 0; i < 100; i++ {
		mid := (low + high) / 2.0
		if f(mid) < 100 {
			low = mid
		} else {
			high = mid
		}
	}

	puts(low)
}
