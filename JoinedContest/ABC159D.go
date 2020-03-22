package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cmb(x int) int {
	return (x * (x - 1)) / 2
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := make([]int, n)
	m := map[int]int{}
	for i := range a {
		a[i] = nextInt()
		m[a[i]]++
	}

	sum := 0
	for _, v := range m {
		sum += cmb(v)
	}

	for i := range a {
		ans := sum
		ans = ans - cmb(m[a[i]]) + cmb(m[a[i]]-1)
		puts(ans)
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
