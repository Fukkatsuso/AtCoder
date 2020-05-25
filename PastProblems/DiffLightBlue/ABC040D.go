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

type UnionFind struct {
	Par  []int
	Rank []int
	Size []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		Par:  make([]int, n),
		Rank: make([]int, n),
		Size: make([]int, n),
	}
	uf.Init(n)
	return uf
}

func (uf *UnionFind) Init(n int) {
	for i := 0; i < n; i++ {
		uf.Par[i] = i
		uf.Rank[i] = 0
		uf.Size[i] = 1
	}
}

func (uf *UnionFind) Find(x int) int {
	for uf.Par[x] != x {
		x, uf.Par[x] = uf.Par[x], uf.Par[uf.Par[x]]
	}
	return x
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.Rank[x] < uf.Rank[y] {
		uf.Par[x] = y
		uf.Size[y] += uf.Size[x]
	} else {
		uf.Par[y] = x
		if uf.Rank[x] == uf.Rank[y] {
			uf.Rank[x]++
		}
		uf.Size[x] += uf.Size[y]
	}
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

type Edge struct {
	a, b, year int
}

type Edges []Edge

func (a Edges) Len() int           { return len(a) }
func (a Edges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool { return a[i].year > a[j].year }

type Query struct {
	id, v, w int
}

type Queries []Query

func (a Queries) Len() int           { return len(a) }
func (a Queries) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Queries) Less(i, j int) bool { return a[i].w > a[j].w }

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	edges := make(Edges, m)
	for i := range edges {
		a, b, y := nextInt()-1, nextInt()-1, nextInt()
		edges[i] = Edge{a, b, y}
	}
	q := nextInt()
	queries := make(Queries, q)
	for i := range queries {
		v, w := nextInt()-1, nextInt()
		queries[i] = Query{i, v, w}
	}
	sort.Sort(edges)
	sort.Sort(queries)

	uf := NewUnionFind(n)
	ans := make([]int, q)
	for i, j := 0, 0; j < q; {
		if i < m {
			edge, query := edges[i], queries[j]
			if edge.year > query.w {
				uf.Unite(edge.a, edge.b)
				i++
			} else {
				ans[query.id] = uf.Size[uf.Find(query.v)]
				j++
			}
		} else {
			query := queries[j]
			ans[query.id] = uf.Size[uf.Find(query.v)]
			j++
		}
	}

	for i := range ans {
		puts(ans[i])
	}
}
