// 解説AC

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

	n, k := getInt(), getInt()

	dp := make([][]int, 2*k)
	for i := range dp {
		dp[i] = make([]int, 2*k)
	}
	for i := 0; i < n; i++ {
		x, y, c := getInt(), getInt(), gets()
		if c == "W" {
			y += k
		}
		x, y = x%(2*k), y%(2*k)
		dp[x][y]++
	}
	// j方向の累積和
	for i := 0; i < 2*k; i++ {
		for j := 1; j < 2*k; j++ {
			dp[i][j] += dp[i][j-1]
		}
	}
	// i方向の累積和
	for j := 0; j < 2*k; j++ {
		for i := 1; i < 2*k; i++ {
			dp[i][j] += dp[i-1][j]
		}
	}

	sum := func(xmin, xmax, ymin, ymax int) int {
		if xmax < 0 || ymax < 0 {
			return 0
		}
		res := dp[xmax][ymax]
		if xmin > 0 {
			res -= dp[xmin-1][ymax]
		}
		if ymin > 0 {
			res -= dp[xmax][ymin-1]
		}
		if xmin > 0 && ymin > 0 {
			res += dp[xmin-1][ymin-1]
		}
		return res
	}

	ans := 0
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			// メイン
			cnt := sum(i, i+k-1, j, j+k-1)
			// 左下
			cnt += sum(0, i-1, 0, j-1)
			// 左上
			cnt += sum(0, i-1, j+k, 2*k-1)
			// 右上
			cnt += sum(i+k, 2*k-1, j+k, 2*k-1)
			// 右下
			cnt += sum(i+k, 2*k-1, 0, j-1)
			ans = max(ans, cnt, n-cnt)
		}
	}
	puts(ans)
}
