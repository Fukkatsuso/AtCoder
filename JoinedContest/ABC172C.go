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

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
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

// keyを超えない最大のインデックス
func binSearch(a []int, key int) int {
	low, high := 0, len(a)
	for low < high-1 {
		mid := (low + high) / 2
		if a[mid] > key {
			high = mid
		} else {
			low = mid
		}
	}
	return low
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

	n, m, k := nextInt3()
	a := nextInts(n)
	b := nextInts(m)

	sa, sb := make([]int, n+1), make([]int, m+1)
	for i := 1; i <= n; i++ {
		sa[i] = sa[i-1] + a[i-1]
	}
	for i := 1; i <= m; i++ {
		sb[i] = sb[i-1] + b[i-1]
	}

	ans := 0
	for i := 0; i <= n; i++ {
		rem := k - sa[i]
		if rem == 0 {
			ans = max(ans, i)
			continue
		} else if rem < 0 {
			continue
		}
		idx := binSearch(sb, rem)
		ans = max(ans, i+idx)
	}

	puts(ans)
}
