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

func bfs(s []string, cost [][]int, h, w, x, y int) {
	ways := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, way := range ways {
		nextx, nexty := x+way[0], y+way[1]
		// 進める場合
		if (0 <= nextx && nextx < w && 0 <= nexty && nexty < h) && s[y][x] == '.' {
			if cost[y][x]+1 < cost[nexty][nextx] {
				cost[nexty][nextx] = cost[y][x] + 1
				bfs(s, cost, h, w, nextx, nexty)
			}
		}
	}
}

func fillSlice(s [][]int, n int) {
	for i := range s {
		for j := range s[i] {
			s[i][j] = n
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := range s {
		s[i] = next()
	}

	ans := 0
	cost := make([][]int, h)
	for i := range cost {
		cost[i] = make([]int, w)
	}

	for i := 0; i < h*w; i++ {
		fillSlice(cost, h*w)
		cost[i/w][i%w] = 0
		bfs(s, cost, h, w, i%w, i/w)
		for j := 0; j < h*w; j++ {
			if i == j || s[i/w][i%w] == '#' || s[j/w][j%w] == '#' {
				continue
			}
			ans = max(ans, cost[j/w][j%w])
		}
	}

	puts(ans)
}
