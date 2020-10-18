// 参考
// https://kakedashi-engineer.appspot.com/2020/05/21/dpl2a/

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
	inf            = 1 << 60
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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x, y, z int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	p := make([]Point, n)
	for i := range p {
		x, y, z := nextInt3()
		p[i] = Point{x, y, z}
	}

	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cost[i][j] = abs(p[j].x-p[i].x) + abs(p[j].y-p[i].y) + max(0, p[j].z-p[i].z)
		}
	}

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(s, cur int) int
	dfs = func(s, cur int) int {
		if dp[s][cur] != -1 {
			return dp[s][cur]
		}
		if s == (1<<n)-1 && cur == 0 {
			return 0
		}

		res := inf
		for i := 0; i < n; i++ {
			// 未訪問の場合
			if s&(1<<i) == 0 {
				res = min(res, dfs(s|(1<<i), i)+cost[cur][i])
			}
		}
		dp[s][cur] = res
		return res
	}

	ans := dfs(0, 0)
	puts(ans)
}
