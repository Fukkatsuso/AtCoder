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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInt2() (int, int) {
	return nextInt(), nextInt()
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

	n, m := nextInt2()

	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			if i == j {
				cost[i][j] = 0
			} else {
				cost[i][j] = inf
			}
		}
	}
	// 0に隣接する頂点のリスト
	neighbor0 := make([]int, 0)
	for i := 0; i < m; i++ {
		u, v, l := nextInt()-1, nextInt()-1, nextInt()
		cost[u][v], cost[v][u] = l, l
		if u == 0 {
			neighbor0 = append(neighbor0, v)
		}
	}
	if len(neighbor0) < 2 {
		puts(-1)
		return
	}

	// ワーシャルフロイド
	for k := 1; k < n; k++ {
		for i := 1; i < n; i++ {
			for j := 1; j < n; j++ {
				cost[i][j] = min(cost[i][j], cost[i][k]+cost[k][j])
			}
		}
	}

	ans := inf
	// 0に隣接する2頂点との距離と，2頂点間の最短距離の合算
	for i := 0; i < len(neighbor0)-1; i++ {
		for j := i + 1; j < len(neighbor0); j++ {
			a, b := neighbor0[i], neighbor0[j]
			ans = min(ans, cost[0][a]+cost[0][b]+cost[a][b])
		}
	}

	if ans == inf {
		puts(-1)
	} else {
		puts(ans)
	}
}
