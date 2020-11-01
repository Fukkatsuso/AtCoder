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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
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

	n, m := nextInt(), nextInt()
	h := nextInts(n)
	w := nextInts(m)
	sort.Sort(sort.IntSlice(h))

	// d1[i=2k]: h[2k]-h[2k-1]
	// d2[i=2k]: h[2k+1]-h[2k]
	d1, d2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i += 2 {
		if i-1 >= 0 {
			d1[i] = d1[i-2] + h[i] - h[i-1]
		}
		if i+1 < n {
			d2[i] = h[i+1] - h[i]
		}
		if i-2 >= 0 {
			d2[i] += d2[i-2]
		}
	}

	ans := 1 << 60
	for i := range w {
		idx := upperBound(h, w[i])
		res := 0
		if idx%2 == 0 {
			// (w[i], h[idx])でペアを組む
			if idx-2 >= 0 {
				res += d2[idx-2]
			}
			if idx < n {
				res += (h[idx] - w[i]) + (d1[n-1] - d1[idx])
			}
		} else {
			// (h[idx-1], w[i])でペアを組む
			if idx-3 >= 0 {
				res += d2[idx-3]
			}
			res += (w[i] - h[idx-1])
			if idx < n {
				res += (d1[n-1] - d1[idx-1])
			}
		}
		ans = min(ans, res)
	}

	puts(ans)
}
