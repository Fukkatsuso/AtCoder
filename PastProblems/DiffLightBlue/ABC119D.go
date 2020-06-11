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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1 << 60

	a, b, q := nextInt(), nextInt(), nextInt()
	s, t := make([]int, a+2), make([]int, b+2)
	s[0], s[a+1], t[0], t[b+1] = -inf, inf, -inf, inf
	for i := 1; i <= a; i++ {
		s[i] = nextInt()
	}
	for i := 1; i <= b; i++ {
		t[i] = nextInt()
	}

	for i := 0; i < q; i++ {
		x := nextInt()
		sr, tr := lowerBound(s, x), lowerBound(t, x)
		sl, tl := sr, tr
		if s[sl] != x {
			sl--
		}
		if t[tl] != x {
			tl--
		}
		ans := min(
			max(x-s[sl], x-t[tl]),
			max(s[sr]-x, t[tr]-x),
			max(0, 2*(x-s[sl])+t[tr]-x),
			max(0, 2*(t[tr]-x)+x-s[sl]),
			max(0, 2*(s[sr]-x)+x-t[tl]),
			max(0, 2*(x-t[tl])+s[sr]-x),
		)
		puts(ans)
	}
}
