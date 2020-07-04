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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

type Edge struct {
	to, cap, rev int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Used  []bool
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
func (g *Graph) AddEdge(from, to, cap int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cap, len(g.Edges[to])})
	g.Edges[to] = append(g.Edges[to], Edge{from, 0, len(g.Edges[from]) - 1})
}

func (g *Graph) dfs(v, t, f int, used []bool) int {
	if v == t {
		return f
	}
	used[v] = true

	for i, e := range g.Edges[v] {
		if used[e.to] || e.cap <= 0 {
			continue
		}
		d := g.dfs(e.to, t, min(f, e.cap), used)
		if d > 0 {
			g.Edges[v][i].cap -= d
			g.Edges[e.to][e.rev].cap += d
			return d
		}
	}
	return 0
}

func (g *Graph) MaxFlow(s, t int) int {
	flow := 0
	const inf = 1 << 60
	for {
		f := g.dfs(s, t, inf, make([]bool, len(g.Edges)+1))
		if f == 0 {
			return flow
		}
		flow += f
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, G, E := nextInt(), nextInt(), nextInt()
	p := nextInts(G)
	g := NewGraph(N + 1)
	for i := 0; i < E; i++ {
		a, b := nextInt(), nextInt()
		g.AddEdge(a, b, 1)
		g.AddEdge(b, a, 1)
	}

	for i := range p {
		g.AddEdge(p[i], N, 1)
	}

	ans := g.MaxFlow(0, N)
	puts(ans)
}
