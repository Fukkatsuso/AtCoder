// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

type Point struct {
	id, x, y int
}

type UnionFind struct {
	Par []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		Par: make([]int, n),
	}
	uf.Init(n)
	return uf
}

func (uf *UnionFind) Init(n int) {
	for i := 0; i < n; i++ {
		uf.Par[i] = -1
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Par[x] < 0 {
		return x
	}
	uf.Par[x] = uf.Find(uf.Par[x])
	return uf.Par[x]
}

func (uf *UnionFind) Unite(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x == y {
		return
	}

	if uf.Par[x] > uf.Par[y] {
		x, y = y, x
	}
	uf.Par[x] += uf.Par[y]
	uf.Par[y] = x
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// xの属する集合の要素数
func (uf *UnionFind) Size(x int) int {
	return -uf.Par[uf.Find(x)]
}

type Edge struct {
	from, to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges []Edge
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges: make([]Edge, 0),
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges = append(g.Edges, Edge{from, to, cost})
}

func (g *Graph) Kruskal(n int) int {
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].cost < g.Edges[j].cost
	})
	uf := NewUnionFind(n)
	cost := 0
	for _, e := range g.Edges {
		if !uf.Same(e.from, e.to) {
			cost += e.cost
			uf.Unite(e.from, e.to)
		}
	}
	return cost
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	p := make([]Point, n)
	for i := 0; i < n; i++ {
		p[i].id = i
		p[i].x, p[i].y = getInt(), getInt()
	}

	g := NewGraph(n)

	// xの昇順
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})
	for i := 0; i < n-1; i++ {
		cost := min(abs(p[i].x-p[i+1].x), abs(p[i].y-p[i+1].y))
		g.AddEdge(p[i].id, p[i+1].id, cost)
		g.AddEdge(p[i+1].id, p[i].id, cost)
	}

	// yの昇順
	sort.Slice(p, func(i, j int) bool {
		return p[i].y < p[j].y
	})
	for i := 0; i < n-1; i++ {
		cost := min(abs(p[i].x-p[i+1].x), abs(p[i].y-p[i+1].y))
		g.AddEdge(p[i].id, p[i+1].id, cost)
		g.AddEdge(p[i+1].id, p[i].id, cost)
	}

	ans := g.Kruskal(n)
	puts(ans)
}
