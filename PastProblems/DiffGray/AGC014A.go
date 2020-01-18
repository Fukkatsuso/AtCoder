// 1つWAでつまづいた
// -1 はabc全て等しくかつ偶数のときのみ

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	defer wt.Flush()
	a, b, c := nextInt(), nextInt(), nextInt()
	if a == b && b == c && (a&1) == 0 {
		puts(-1)
		return
	}
	ans := 0
	for ((a & 1) | (b & 1) | (c & 1)) == 0 {
		a, b, c = (b+c)/2, (c+a)/2, (a+b)/2
		ans++
	}
	puts(ans)
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
