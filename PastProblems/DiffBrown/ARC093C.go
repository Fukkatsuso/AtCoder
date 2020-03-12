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

	n := nextInt()
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}

	s := 0
	for i := 1; i < n+2; i++ {
		s += abs(a[i] - a[i-1])
	}

	for i := 1; i <= n; i++ {
		puts(s + abs(a[i+1]-a[i-1]) - abs(a[i]-a[i-1]) - abs(a[i+1]-a[i]))
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

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func nextFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = nextFloat64()
	}
	return slice
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
