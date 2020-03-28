// コンテスト後AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, y, a, b, c := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	p, q, r := nextInts(a), nextInts(b), nextInts(c)
	sort.Sort(sort.IntSlice(p))
	sort.Sort(sort.IntSlice(q))
	sort.Sort(sort.IntSlice(r))

	p = p[a-x:]
	q = q[b-y:]

	for i, j, k := 0, 0, c-1; (i < x || j < y) && k >= 0; k-- {
		if i < x && j < y && p[i] < q[j] && p[i] < r[k] {
			// pの要素が最小
			p[i] = r[k]
			i++
		} else if i < x && j < y && p[i] > q[j] && q[j] < r[k] {
			// qの要素が最小
			q[j] = r[k]
			j++
		} else if i < x && p[i] < r[k] {
			p[i] = r[k]
			i++
		} else if j < y && q[j] < r[k] {
			q[j] = r[k]
			j++
		} else {
			break
		}
	}

	ans := sum(p) + sum(q)
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
