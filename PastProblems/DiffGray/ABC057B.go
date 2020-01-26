package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	a, b := make([]int, n), make([]int, n)
	c, d := make([]int, m), make([]int, m)

	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	for i := 0; i < m; i++ {
		c[i], d[i] = nextInt(), nextInt()
	}

	for i := 0; i < n; i++ {
		min := dist(a[i], b[i], c[0], d[0])
		minJ := 0
		for j := 1; j < m; j++ {
			if d := dist(a[i], b[i], c[j], d[j]); d < min {
				min = d
				minJ = j
			}
		}
		puts(minJ + 1)
	}
	wt.Flush()
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
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

func nextMega() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
