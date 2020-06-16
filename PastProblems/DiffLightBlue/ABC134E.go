// 解説AC

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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func binarySearch(a []int, key int) int {
	l, r := -1, len(a)-1
	for r-l > 1 {
		mid := (l + r) / 2
		if a[mid] < key {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)

	colors := make([]int, n)
	for i := range colors {
		colors[i] = -1
	}
	for i := range a {
		x := binarySearch(colors, a[i])
		colors[x] = a[i]
	}

	ans := 0
	for i := range colors {
		if colors[i] < 0 {
			break
		}
		ans++
	}

	puts(ans)
}
