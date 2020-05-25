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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
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

	h, w := nextInt2()
	a := make([][]int, h)
	for i := range a {
		a[i] = nextInts(w)
	}

	g := NewGraph(h * w)
	di := []int{0, 1, 0, -1}
	dj := []int{1, 0, -1, 0}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 4; k++ {
				nextI, nextJ := i+di[k], j+dj[k]
				if nextI < 0 || h <= nextI || nextJ < 0 || w <= nextJ {
					continue
				}
				if a[i][j] < a[nextI][nextJ] {
					g.AddEdge(i*w+j, nextI*w+nextJ, 1)
				}
			}
		}
	}

	dp := make([]int, h*w)
	var dfs func(par, cur int)
	dfs = func(par, cur int) {
		for _, e := range g.Edges[cur] {
			if dp[e.to] == 0 {
				dfs(cur, e.to)
			}
		}
		dp[cur] = 1
		for _, e := range g.Edges[cur] {
			dp[cur] = (dp[cur] + dp[e.to]) % mod
		}
	}
	for i := range dp {
		if dp[i] == 0 {
			dfs(-1, i)
		}
	}

	ans := 0
	for i := range dp {
		ans = (ans + max(1, dp[i])) % mod
	}

	puts(ans)
}
