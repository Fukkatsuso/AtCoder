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

func dfs(s []string, cost [][]int, h, w, x, y int) {
	ways := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, way := range ways {
		nextX, nextY := x+way[0], y+way[1]
		if nextX < 0 || w <= nextX || nextY < 0 || h <= nextY {
			continue
		}
		if s[nextY][nextX] == '.' && cost[y][x]+1 < cost[nextY][nextX] {
			cost[nextY][nextX] = cost[y][x] + 1
			dfs(s, cost, h, w, nextX, nextY)
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

	cost := make([][]int, h)
	for i := range cost {
		cost[i] = make([]int, w)
		for j := range cost[i] {
			cost[i][j] = 2 * h * w
		}
	}
	cost[0][0] = 0
	dfs(s, cost, h, w, 0, 0)

	white := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				white++
			}
		}
	}

	if cost[h-1][w-1] > h*w {
		puts(-1)
	} else {
		puts(white - 1 - cost[h-1][w-1])
	}
}
