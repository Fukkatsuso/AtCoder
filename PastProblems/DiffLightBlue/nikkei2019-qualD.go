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

type Edge struct {
	to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges [][]Edge // Edge[i][j]: 頂点iのj番目の辺
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges: make([][]Edge, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost})
}

func topologicalSort(n int, g *Graph, indegree []int) []int {
	sorted := make([]int, 0)

	que := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		for _, e := range g.Edges[v] {
			u := e.to
			indegree[u]--
			if indegree[u] == 0 {
				que = append(que, u)
			}
		}
		sorted = append(sorted, v)
	}
	return sorted
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	g := NewGraph(n)
	indegree := make([]int, n) // 入次数
	for i := 0; i < n-1+m; i++ {
		a, b := nextInt()-1, nextInt()-1
		g.AddEdge(a, b, 1)
		indegree[b]++
	}

	sorted := topologicalSort(n, g, indegree)
	// mp[v]: 頂点vのソート後のインデックス
	mp := map[int]int{}
	for i, v := range sorted {
		mp[v] = i
	}
	mp[-1] = -1

	par := make([]int, n)
	for i := range par {
		par[i] = -1
	}
	for v, edges := range g.Edges {
		for _, e := range edges {
			if mp[v] > mp[par[e.to]] {
				par[e.to] = v
			}
		}
	}
	for _, p := range par {
		puts(p + 1)
	}
}
