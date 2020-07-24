// 再帰がうまく書けずWA

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

func (g *Graph) dfs(prev, cur int, h []int) int {
	cost := 0
	for _, e := range g.Edges[cur] {
		if e.to == prev {
			continue
		}
		cost += g.dfs(cur, e.to, h)
	}
	if prev == -1 {
		// root
		return cost
	} else if cost == 0 {
		// 葉
		return h[cur]
	}
	return cost + 1
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, x := nextInt(), nextInt()-1
	h := nextInts(n)
	g := NewGraph(n)
	for i := 0; i < n-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		g.AddEdge(a, b, 1)
		g.AddEdge(b, a, 1)
	}

	ans := g.dfs(-1, x, h) * 2
	puts(ans)
}
