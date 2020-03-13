package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type String []byte

func (s String) Len() int {
	return len(s)
}

func (s String) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s String) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	m := map[string]int{}
	ans := 0
	for i := 0; i < n; i++ {
		s := String(next())
		sort.Sort(s)
		ans += m[string(s)]
		m[string(s)]++
	}

	puts(ans)
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
