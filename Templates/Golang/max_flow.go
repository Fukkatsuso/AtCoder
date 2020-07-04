package main

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

// Ford-Fulkerson法
// 使い方はABC010D（青diff）参考
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
