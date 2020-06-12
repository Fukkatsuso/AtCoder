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
	Edges    [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Pred     []int    // 遷移元
	Cost     []int    // 始点からの最小コスト
	RouteNum []int    // 始点から各頂点への最短経路数
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
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
	return g
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
		u, v, w := nextInt()-1, nextInt()-1, nextInt()
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	color := make([]int, n)
	color[0] = 1
	for i := 1; i < n; i++ {
		color[i] = -1
	}

	q := make([]int, 0)
	q = append(q, 0)
	for len(q) > 0 {
		cur := q[0]
		for _, e := range g.Edges[cur] {
			if color[e.to] < 0 {
				if e.cost%2 == 0 {
					color[e.to] = color[cur]
				} else {
					color[e.to] = 1 - color[cur]
				}
				q = append(q, e.to)
			}
		}
		q = q[1:]
	}

	for _, c := range color {
		puts(c)
	}
}
