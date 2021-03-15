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
	inf            = 1 << 60
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

	n, m, l := getInt(), getInt(), getInt()

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}
	for i := 0; i < m; i++ {
		a, b, c := getInt()-1, getInt()-1, getInt()
		dist[a][b], dist[b][a] = c, c
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost {
			if dist[i][j] <= l {
				cost[i][j] = 1
			} else {
				cost[i][j] = inf
			}
		}
		cost[i][i] = 0
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				cost[i][j] = min(cost[i][j], cost[i][k]+cost[k][j])
			}
		}
	}

	q := getInt()
	for i := 0; i < q; i++ {
		s, t := getInt()-1, getInt()-1
		ans := cost[s][t]
		if ans == inf {
			ans = 0
		}
		puts(ans - 1)
	}
}
