// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	l := nextInts(n)
	sort.Sort(sort.IntSlice(l))

	ans := 0
	for i := n - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			a, b := l[i], l[j]
			lb := lowerBound(l, a-b+1)
			// puts(lb, j-lb)
			ans += max(0, j-lb)
		}
	}
	puts(ans)
}

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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func lowerBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := l + (r-l)/2
		if a[mid] < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
