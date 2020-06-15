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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Edge struct {
	to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges    [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Cost     [][3]int // 始点からの最小コスト
	RouteNum []int    // 始点から各頂点への最短経路数
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges:    make([][]Edge, n),
		Cost:     make([][3]int, n),
		RouteNum: make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
		g.Cost[i] = [3]int{-1, -1, -1}
		g.RouteNum[i] = 0
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost})
}

func bfs(g *Graph, s, t int) {
	que := make([][2]int, 0)
	g.Cost[s][0] = 0
	que = append(que, [2]int{s, 0})
	for len(que) > 0 {
		cur, parity := que[0][0], que[0][1]
		for _, e := range g.Edges[cur] {
			nextParity := (parity + 1) % 3
			if g.Cost[e.to][nextParity] == -1 {
				g.Cost[e.to][nextParity] = g.Cost[cur][parity] + 1
				que = append(que, [2]int{e.to, nextParity})
			}
		}
		que = que[1:]
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		u, v := nextInt()-1, nextInt()-1
		g.AddEdge(u, v, 1)
	}
	s, t := nextInt()-1, nextInt()-1

	bfs(g, s, t)

	if g.Cost[t][0] == -1 {
		puts(-1)
	} else {
		puts(g.Cost[t][0] / 3)
	}
}
