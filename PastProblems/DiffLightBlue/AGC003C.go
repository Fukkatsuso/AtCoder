package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for l, r := 0, len(a); l < r; {
		mid := (l + r) / 2
		if a[mid] == key {
			return mid
		}
		if key < a[mid] {
			r = mid
		} else if a[mid] < key {
			l = mid + 1
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		b[i] = a[i]
	}

	sort.Sort(sort.IntSlice(b))

	ans := 0
	for i := 0; i < n; i++ {
		idx := binarySearch(b, a[i])
		ans += abs(i-idx) % 2
	}

	puts(ans / 2)
}
