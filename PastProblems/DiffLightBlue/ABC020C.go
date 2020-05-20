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

func dfs(s []string, h, w, si, sj, x int, cost [][]int) {
	di := []int{0, 1, 0, -1}
	dj := []int{1, 0, -1, 0}
	var rec func(i, j int)
	rec = func(i, j int) {
		if s[i][j] == 'G' {
			return
		}
		for k := 0; k < 4; k++ {
			nextI, nextJ := i+di[k], j+dj[k]
			if nextI < 0 || h <= nextI || nextJ < 0 || w <= nextJ {
				continue
			}
			if (s[nextI][nextJ] == '.' || s[nextI][nextJ] == 'G') && cost[nextI][nextJ] > cost[i][j]+1 {
				cost[nextI][nextJ] = cost[i][j] + 1
				rec(nextI, nextJ)
			} else if s[nextI][nextJ] == '#' && cost[nextI][nextJ] > cost[i][j]+x {
				cost[nextI][nextJ] = cost[i][j] + x
				rec(nextI, nextJ)
			}
		}
	}
	rec(si, sj)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1 << 60

	h, w, t := nextInt3()
	s := make([]string, h)
	var si, sj, gi, gj int
	for i := 0; i < h; i++ {
		s[i] = next()
		for j := 0; j < w; j++ {
			if s[i][j] == 'S' {
				si, sj = i, j
			} else if s[i][j] == 'G' {
				gi, gj = i, j
			}
		}
	}

	cost := make([][]int, h)
	for i := range cost {
		cost[i] = make([]int, w)
	}
	low := 1
	for high := t + 1; low < high-1; {
		// init cost
		for i := range cost {
			for j := range cost[i] {
				cost[i][j] = inf
			}
		}
		cost[si][sj] = 0
		// search
		mid := (low + high) / 2
		dfs(s, h, w, si, sj, mid, cost)
		// update
		if cost[gi][gj] <= t {
			low = mid
		} else {
			high = mid
		}
	}

	puts(low)
}
