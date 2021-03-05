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

func sum(a ...int) int {
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

	n, k := getInt(), getInt()
	a := getInts(n)

	ok := func(x int) bool {
		// a[i]%xを昇順に並べる
		r := make([]int, n)
		for i := 0; i < n; i++ {
			r[i] = a[i] % x
		}
		sort.Ints(r)
		// rの累積和, rsum[i]: \sigma_{j=0}^{j=i-1}(r[j])
		rsum := make([]int, n+1)
		for i := 0; i < n; i++ {
			rsum[i+1] = r[i] + rsum[i]
		}
		for i := 0; i <= n && rsum[i] <= k; i++ {
			// r[:i]だけ引き，\sigma(x-r[j])(i<=j<n)だけ加算
			if rsum[i] == (n-i)*x-(rsum[n]-rsum[i]) {
				return true
			}
		}
		return false
	}

	ans := 1
	s := sum(a...)
	for i := 1; i*i <= s; i++ {
		if s%i != 0 {
			continue
		}
		if i > ans && ok(i) {
			ans = i
		}
		if s/i > ans && ok(s/i) {
			ans = s / i
		}
	}

	puts(ans)
}
