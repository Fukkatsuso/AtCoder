package main

import (
	"bufio"
	"fmt"
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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	a := getInts(n)

	asum := make([]int, n)
	aasum := make([]int, n)
	for i := 0; i < n; i++ {
		asum[i] = a[i]
		aasum[i] = a[i] * a[i]
		if i > 0 {
			asum[i] += asum[i-1]
			aasum[i] += aasum[i-1]
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		add := a[i]*a[i]*i - 2*a[i]*asum[i-1] + aasum[i-1]
		ans += add
	}
	puts(ans)
}
