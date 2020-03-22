package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverseString(s string) string {
	n := len(s)
	bytes := []byte(s)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}
	return string(bytes)
}

func kaibun(s string) bool {
	return s == reverseString(s)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	n := len(s)

	ok := kaibun(s)
	ok = ok && kaibun(s[0:(n-1)/2])
	ok = ok && kaibun(s[(n+3)/2-1:n])

	if ok {
		puts("Yes")
	} else {
		puts("No")
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
