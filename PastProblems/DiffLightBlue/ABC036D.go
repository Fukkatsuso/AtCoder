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
	mod            = 1000000007
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

type Edge struct {
	to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges    [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Pred     []int    // 遷移元
	Cost     []int    // 始点からの最小コスト
	RouteNum []int    // 始点から各頂点への最短経路数
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := Graph{
		Edges:    make([][]Edge, n),
		Pred:     make([]int, n),
		Cost:     make([]int, n),
		RouteNum: make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
		g.Pred[i] = -1
		g.Cost[i] = 1 << 60
		g.RouteNum[i] = 0
	}
	return &g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost})
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	g := NewGraph(n)
	for i := 0; i < n-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		g.AddEdge(a, b, 1)
		g.AddEdge(b, a, 1)
	}

	// dp[i]: f(i), g(i)
	dp := make([][2]int, n)

	var dfs func(par, cur int)
	dfs = func(par, cur int) {
		for _, e := range g.Edges[cur] {
			if e.to != par {
				dfs(cur, e.to)
			}
		}
		dp[cur][0], dp[cur][1] = 1, 1
		for _, e := range g.Edges[cur] {
			if e.to != par {
				dp[cur][0] = (dp[cur][0] * dp[e.to][1]) % mod
				dp[cur][1] = (dp[cur][1] * dp[e.to][0]) % mod
			}
		}
		dp[cur][0] = (dp[cur][0] + dp[cur][1]) % mod
	}

	dfs(-1, 0)

	puts(dp[0][0])
}
