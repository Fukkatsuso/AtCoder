package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isEvenString(s string) bool {
	if len(s)&1 == 1 {
		return false
	}
	if s[0:len(s)/2] == s[len(s)/2:len(s)] {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	s := next()
	for ans := len(s) - 1; ans >= 1; ans-- {
		if isEvenString(s[0:ans]) {
			puts(ans)
			break
		}
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