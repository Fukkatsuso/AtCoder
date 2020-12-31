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
	inf            = 1 << 50
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	need := map[int]int{
		1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6,
	}

	n, m := getInt(), getInt()
	a := getInts(m)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -inf
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := range a {
			if i-need[a[j]] >= 0 {
				dp[i] = max(dp[i], dp[i-need[a[j]]]+1)
			}
		}
	}

	rem := n
	ans := []byte{}
	for i := dp[n]; i > 0; i-- {
		for _, x := range a {
			if rem >= need[x] && dp[rem-need[x]] == dp[rem]-1 {
				ans = append(ans, byte('0'+x))
				rem -= need[x]
				break
			}
		}
	}
	puts(string(ans))
}
