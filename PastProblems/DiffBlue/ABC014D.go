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

// https://algo-logic.info/lca/
type LCA struct {
	Parent [][]int
	Dist   []int
}

// LCA: n頂点の木の最小共通祖先
// O(NlogN)
func GetLCA(g *Graph, n, root int) *LCA {
	// init
	k := 1
	for (1 << k) < n {
		k++
	}
	lca := LCA{Parent: make([][]int, k), Dist: make([]int, n)}
	for i := range lca.Parent {
		lca.Parent[i] = make([]int, n)
		for j := range lca.Parent[i] {
			lca.Parent[i][j] = -1
		}
	}
	for i := range lca.Dist {
		lca.Dist[i] = -1
	}
	// dfsで各頂点についてrootからの距離と親を求める
	var dfs func(cur, par, dist int)
	dfs = func(cur, par, dist int) {
		lca.Parent[0][cur] = par
		lca.Dist[cur] = dist
		for _, e := range g.Edges[cur] {
			if e.to != par {
				dfs(e.to, cur, dist+1)
			}
		}
	}
	dfs(root, -1, 0)
	// ダブリング
	for i := 0; i < k-1; i++ {
		for v := 0; v < n; v++ {
			if lca.Parent[i][v] < 0 {
				lca.Parent[i+1][v] = -1
			} else {
				lca.Parent[i+1][v] = lca.Parent[i][lca.Parent[i][v]]
			}
		}
	}
	return &lca
}

// 頂点u,vの最小共通祖先を求める
// O(logN)
func (lca *LCA) Query(u, v int) int {
	if lca.Dist[u] < lca.Dist[v] {
		u, v = v, u
	}
	k := len(lca.Parent)
	for i := 0; i < k; i++ {
		if (lca.Dist[u]-lca.Dist[v])&(1<<i) > 0 {
			u = lca.Parent[i][u]
		}
	}
	if u == v {
		return u
	}
	for i := k - 1; i >= 0; i-- {
		if lca.Parent[i][u] != lca.Parent[i][v] {
			u, v = lca.Parent[i][u], lca.Parent[i][v]
		}
	}
	return lca.Parent[0][u]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	g := NewGraph(n)
	for i := 0; i < n-1; i++ {
		x, y := nextInt()-1, nextInt()-1
		g.AddEdge(x, y, 1)
		g.AddEdge(y, x, 1)
	}

	lca := GetLCA(g, n, 0)
	q := nextInt()
	for i := 0; i < q; i++ {
		a, b := nextInt()-1, nextInt()-1
		r := lca.Query(a, b)
		ans := lca.Dist[a] + lca.Dist[b] - 2*lca.Dist[r] + 1
		puts(ans)
	}
}
