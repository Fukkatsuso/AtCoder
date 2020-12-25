// 解説AC?
// 二分探索のバグが取れず

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	a := getInts(n)

	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1] = a[i]
		if i > 0 {
			dp[i+1] += dp[i]
		}
	}

	// C,D,Eの先頭を決めたときのP,Q,R,Sの最大値と最小値の差の絶対値
	f := func(ci, di, ei int) int {
		p, q, r, s := dp[ci], dp[di]-dp[ci], dp[ei]-dp[di], dp[n]-dp[ei]
		return max(p, q, r, s) - min(p, q, r, s)
	}

	ans := 1 << 60
	// di: 部分列Dの先頭
	for di := 2; di < n-1; di++ {
		low, high := 1, di
		for high-low > 1 {
			mid := (low + high) / 2
			if dp[mid] < dp[di]-dp[mid] {
				low = mid
			} else {
				high = mid
			}
		}
		ci := high

		low, high = di+1, n
		for high-low > 1 {
			mid := (low + high) / 2
			if dp[mid]-dp[di] < dp[n]-dp[mid] {
				low = mid
			} else {
				high = mid
			}
		}
		ei := high

		ans = min(ans,
			f(ci, di, ei),
			f(ci-1, di, ei),
			f(ci, di, ei-1),
			f(ci-1, di, ei-1))
	}

	puts(ans)
}
