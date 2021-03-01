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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

// mを10進数からn進数の数に変換
func baseConv(m, n int) []int {
	s := []int{}
	for m > 0 {
		s = append(s, m%n)
		m /= n
	}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}

// x <= y ?
func leqq(x, y []int) bool {
	if len(x) != len(y) {
		return len(x) < len(y)
	}
	n := len(x)
	for i := 0; i < n; i++ {
		if x[i] != y[i] {
			return x[i] < y[i]
		}
	}
	// x == y
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, m := gets(), getInt()
	xnum := make([]int, len(x))
	for i := range x {
		xnum[i] = int(x[i] - '0')
	}

	if len(x) == 1 {
		if xnum[0] <= m {
			puts(1)
		} else {
			puts(0)
		}
		return
	}

	d := 0
	for i := range xnum {
		d = max(d, xnum[i])
	}

	// n > d
	ok, ng := d, 1<<60
	for ng-ok > 1 {
		mid := (ok + ng) / 2
		if leqq(xnum, baseConv(m, mid)) {
			ok = mid
		} else {
			ng = mid
		}
	}

	puts(ok - d)
}
