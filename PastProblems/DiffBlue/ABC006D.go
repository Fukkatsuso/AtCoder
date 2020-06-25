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

func lowerBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func LIS(a []int) int {
	t := make([]int, 0)
	for i := range a {
		idx := lowerBound(t, a[i])
		if idx == len(t) {
			t = append(t, a[i])
		} else {
			t[idx] = a[i]
		}
	}
	return len(t)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	c := nextInts(n)

	ans := n - LIS(c)

	puts(ans)
}
