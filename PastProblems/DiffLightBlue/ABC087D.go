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
	Edges [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Pred  []int    // 遷移元
	Cost  []int    // 始点からの最小コスト
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges: make([][]Edge, n),
		Pred:  make([]int, n),
		Cost:  make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
		g.Pred[i] = -1
		g.Cost[i] = 1 << 60
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

	n, m := nextInt(), nextInt()
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		l, r, d := nextInt()-1, nextInt()-1, nextInt()
		g.AddEdge(l, r, d)
		g.AddEdge(r, l, -d)
	}

	x := make([]int, n)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		q := []int{i}
		for len(q) > 0 {
			v := q[0]
			visited[v] = true
			for _, e := range g.Edges[v] {
				if !visited[e.to] {
					x[e.to] = x[v] + e.cost
					visited[e.to] = true
					q = append(q, e.to)
				} else if x[e.to] != x[v]+e.cost {
					puts("No")
					return
				}
			}
			q = q[1:]
		}
	}
	puts("Yes")
}
