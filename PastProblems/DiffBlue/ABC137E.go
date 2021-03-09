// 解説AC
// 参考: https://sigma1113.hatenablog.com/entry/2019/08/12/130042

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

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
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

type Edge struct {
	to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Cost  []int    // 始点からの最小コスト
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges: make([][]Edge, n),
		Cost:  make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
		g.Cost[i] = inf
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost})
}

// sからtへの経路に負の閉路があるかを返す
func (g *Graph) BellmanFord(n, s, t int) bool {
	g.Cost[s] = 0
	// n-1回の更新
	for i := 0; i < n-1; i++ {
		// 全ての辺を見る
		for v := range g.Edges {
			for _, e := range g.Edges[v] {
				if g.Cost[v] != inf && g.Cost[e.to] > g.Cost[v]+e.cost {
					g.Cost[e.to] = g.Cost[v] + e.cost
				}
			}
		}
	}
	// 閉路検出
	neg := make([]bool, n)
	for i := 0; i < n; i++ {
		for v := range g.Edges {
			for _, e := range g.Edges[v] {
				if g.Cost[e.to] != inf && g.Cost[e.to] > g.Cost[v]+e.cost {
					g.Cost[e.to] = g.Cost[v] + e.cost
					neg[e.to] = true
				}
				if neg[v] {
					neg[e.to] = true
				}
			}
		}
	}
	return neg[t]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, p := getInt3()
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, c := getInt()-1, getInt()-1, getInt()
		cost := -(c - p)
		g.AddEdge(a, b, cost)
	}

	closed := g.BellmanFord(n, 0, n-1)
	ans := max(0, -g.Cost[n-1])
	if closed {
		ans = -1
	}
	puts(ans)
}
