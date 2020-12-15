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

func divCeil(a, b int) int {
	return (a + (b - 1)) / b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = getInt() - 1
	}
	sort.Sort(sort.IntSlice(a))

	if m == 0 {
		puts(1)
		return
	}

	k := n
	for i := 1; i < m; i++ {
		if a[i]-a[i-1] > 1 {
			k = min(k, a[i]-a[i-1]-1)
		}
	}
	// 左端
	if a[0] > 0 {
		k = min(k, a[0])
	}
	// 右端
	if a[m-1] < n-1 {
		k = min(k, n-1-a[m-1])
	}
	k = max(1, k)

	ans := 0
	for i := 1; i < m; i++ {
		ans += divCeil(a[i]-a[i-1]-1, k)
	}
	// 左端
	if a[0] > 0 {
		ans += divCeil(a[0], k)
	}
	// 右端
	if a[m-1] < n-1 {
		ans += divCeil(n-1-a[m-1], k)
	}

	puts(ans)
}
