// 解説AC
// 時刻t-1に美味しさmaxの料理を注文するべき
// foodsを昇順ソートしてなくてWA

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

type Food struct {
	a, b int
}

type Foods []Food

func (a Foods) Len() int      { return len(a) }
func (a Foods) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Foods) Less(i, j int) bool {
	return a[i].a < a[j].a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, t := nextInt(), nextInt()
	foods := make(Foods, n)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		foods[i] = Food{a, b}
	}
	sort.Sort(foods)

	// dp[i+1][j]: 料理iを時刻jに食べ終わるときの満足度の最大値
	dp := make([][10001]int, n+1)
	for i := 0; i < n; i++ {
		for j := 1; j <= 10000; j++ {
			// j-foods[i].a: 料理iを注文して食べ始める時間
			if 0 <= j-foods[i].a && j-foods[i].a < t {
				dp[i+1][j] = max(dp[i][j], dp[i][j-foods[i].a]+foods[i].b)
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	ans := 0
	for i := range dp[n] {
		ans = max(ans, dp[n][i])
	}
	puts(ans)
}
