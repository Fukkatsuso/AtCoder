// 解説AC
// 閉路があれば無条件でinfにしていた

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
	inf            = 1 << 60
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

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Edge struct {
	from, to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges []Edge // Edge[i]: i番目の辺
	Cost  []int  // 始点からの最小コスト
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges: make([]Edge, 0),
		Cost:  make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Cost[i] = inf
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges = append(g.Edges, Edge{from, to, cost})
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, c := getInt()-1, getInt()-1, getInt()
		// コストを-1倍して最短距離問題にする
		g.AddEdge(a, b, -c)
	}

	g.Cost[0] = 0
	for i := 0; i <= n; i++ {
		for _, e := range g.Edges {
			if g.Cost[e.from] != inf && g.Cost[e.to] > g.Cost[e.from]+e.cost {
				g.Cost[e.to] = g.Cost[e.from] + e.cost
				// 頂点nを含む閉路があればinf
				if i == n && e.to == n-1 {
					puts("inf")
					return
				}
			}
		}
	}
	puts(-g.Cost[n-1])
}
