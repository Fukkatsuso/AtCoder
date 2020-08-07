// 解説AC
// 要復習

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

// key以上の要素の最小インデックス
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

// keyより大きい要素の最小インデックス
func upperBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] <= key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n := nextInt()
	d := nextInts(n)
	sort.Sort(sort.IntSlice(d))

	A, B := make([]int, n), make([]int, n)
	for i, x := range d {
		A[i] = n - lowerBound(d, 2*x)
		B[i] = upperBound(d, x/2)
	}

	for i := n - 2; i >= 0; i-- {
		A[i] += A[i+1]
		A[i] %= mod
	}

	ans := 0
	for i, x := range d {
		idx := lowerBound(d, 2*x)
		if idx < n {
			add := (B[i] * A[idx]) % mod
			ans += add
			ans %= mod
		}
	}
	puts(ans)
}
