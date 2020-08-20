// 解説AC
// 特に難しいと感じた

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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

// key以上の要素の最小インデックス
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n := nextInt()
	c := nextInts(n)
	m := max(c...)

	places := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		places[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		places[c[i]] = append(places[c[i]], i)
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1]

		// 同じ色の石が左側にいくつあるか
		color := c[i-1]
		idx := lowerBound(places[color], i-1)
		if idx > 0 {
			// 一番近くにある同色の石の位置(0-indexed)
			j := places[color][idx-1]
			if (i-1)-j > 1 {
				dp[i] += dp[j+1]
			}
		}

		dp[i] %= mod
	}

	puts(dp[n])
}
