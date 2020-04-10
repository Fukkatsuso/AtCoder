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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	r, c := nextInt(), nextInt()
	s, g := []int{nextInt() - 1, nextInt() - 1}, []int{nextInt() - 1, nextInt() - 1}
	maze := make([]string, r)
	for i := range maze {
		maze[i] = next()
	}
	// 手数
	cost := make([][]int, r)
	for i := range cost {
		cost[i] = make([]int, c)
		for j := range cost[i] {
			cost[i][j] = -1
		}
	}
	cost[s[0]][s[1]] = 0

	que := make([][]int, 0)
	empty := func() bool {
		return len(que) == 0
	}
	push := func(p []int) {
		que = append(que, p)
	}
	pop := func() []int {
		p := que[0]
		que = que[1:]
		return p
	}

	push([]int{s[1], s[0]})
	ways := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for !empty() {
		p := pop()
		x, y := p[0], p[1]
		for _, way := range ways {
			nextX, nextY := x+way[0], y+way[1]
			if nextX < 0 || c <= nextX || nextY < 0 || r <= nextY {
				continue
			}
			if maze[nextY][nextX] == '.' && cost[nextY][nextX] < 0 {
				cost[nextY][nextX] = cost[y][x] + 1
				push([]int{nextX, nextY})
			}
		}
	}

	puts(cost[g[0]][g[1]])
}
