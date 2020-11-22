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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

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

	n, t := getInt(), getInt()
	a := getInts(n)

	X := make([]int, 0)
	for bit := 0; bit < (1 << (n / 2)); bit++ {
		sum := 0
		for i := 0; i < (n / 2); i++ {
			if bit&(1<<i) > 0 {
				sum += a[i]
			}
		}
		if sum <= t {
			X = append(X, sum)
		}
	}

	Y := make([]int, 0)
	for bit := 0; bit < (1 << ((n + 1) / 2)); bit++ {
		sum := 0
		for i := 0; i < ((n + 1) / 2); i++ {
			if bit&(1<<i) > 0 {
				sum += a[n/2+i]
			}
		}
		if sum <= t {
			Y = append(Y, sum)
		}
	}
	sort.Sort(sort.IntSlice(Y))

	ans := 0
	for _, x := range X {
		idx := upperBound(Y, t-x) - 1
		if idx >= 0 {
			y := Y[idx]
			ans = max(ans, x+y)
		} else {
			ans = max(ans, x)
		}
	}
	puts(ans)
}
