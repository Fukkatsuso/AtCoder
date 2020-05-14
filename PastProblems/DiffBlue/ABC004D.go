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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	r, g, b := nextInt3()
	sum := r + g + b

	// 中心0を500に変換して計算
	cost := func(pos, rem int) int {
		// red
		if rem >= g+b {
			return abs(pos - 400)
		}
		// green
		if rem >= b {
			return abs(pos - 500)
		}
		// blue
		return abs(pos - 600)
	}

	dp := make([][]int, 1000)
	for i := range dp {
		dp[i] = make([]int, 1000)
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
		dp[i][sum] = 0
	}
	for i := 1; i < 1000; i++ {
		for j := 0; j < sum; j++ {
			// 座標iで残りマーブルj個となる最小のコスト
			// min(
			//	座標iにマーブルを置かない場合
			// 	座標iにマーブルを置く場合(ここに置くことでマーブルがj個になる)
			// )
			dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]+cost(i, j))
		}
	}

	ans := 1 << 60
	for i := 0; i < 1000; i++ {
		ans = min(ans, dp[i][0])
	}

	puts(ans)
}
