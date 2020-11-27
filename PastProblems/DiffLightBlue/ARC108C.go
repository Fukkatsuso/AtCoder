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

type Edge struct {
	to, cost int
}

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

// 辺の削除(成功:true, 失敗:false)
func (g *Graph) DeleteEdge(from, to int) bool {
	for i := range g.Edges[from] {
		if g.Edges[from][i].to == to {
			g.Edges[from] = append(g.Edges[from][:i], g.Edges[from][i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		u, v, c := getInt()-1, getInt()-1, getInt()
		g.AddEdge(u, v, c)
		g.AddEdge(v, u, c)
	}

	// 頂点に付与するラベル
	label := make([]int, n)
	for i := range label {
		label[i] = -1
	}

	q := make([]int, 0)
	q = append(q, 0)
	label[0] = 1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, e := range g.Edges[cur] {
			if label[e.to] > 0 {
				continue
			}
			if e.cost == label[cur] {
				label[e.to] = 1 + (e.cost+1)%n
			} else {
				label[e.to] = e.cost
			}
			q = append(q, e.to)
		}
	}

	for i := range label {
		puts(label[i])
	}
}
