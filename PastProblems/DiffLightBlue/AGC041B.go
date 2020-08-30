// 解説AC

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

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, v, p := nextInt4()
	a := nextInts(n)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	sa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sa[i] = sa[i-1] + a[i-1]
	}

	possible := func(x int) bool {
		if x < p {
			return true
		}
		if a[x]+m < a[p-1] {
			return false
		}
		// can: 問題xを選択可能にするために，入れられても構わない票数
		// 問題x, 上位p-1問, 問題x+1以降
		can := m + m*(p-1) + m*(n-x-1)
		// 問題i:p-1<=i<x
		if p-1 < x {
			can += (a[x]+m)*(x-p+1) - (sa[x] - sa[p-1])
		}
		return can >= m*v
	}

	ok, ng := 0, n
	for ok < ng-1 {
		mid := (ok + ng) / 2
		if possible(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	puts(ok + 1)
}
