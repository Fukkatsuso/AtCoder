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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

type Edge struct {
	to int
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
func (g *Graph) AddEdge(from, to int) {
	g.Edges[from] = append(g.Edges[from], Edge{to})
}

func dfs(g *Graph, a, high []int, cur int) {
	h := -inf
	for _, e := range g.Edges[cur] {
		if high[e.to] == -inf {
			dfs(g, a, high, e.to)
		}
		h = max(h, high[e.to], a[e.to])
	}
	high[cur] = h
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	a := getInts(n)
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		x, y := getInt()-1, getInt()-1
		g.AddEdge(x, y)
	}

	// high[i]: 街iの子孫の最高値
	high := make([]int, n)
	for i := range high {
		high[i] = -inf
	}
	for i := 0; i < n; i++ {
		if high[i] == -inf {
			dfs(g, a, high, i)
		}
	}

	ans := -inf
	for i := 0; i < n; i++ {
		ans = max(ans, high[i]-a[i])
	}
	puts(ans)
}
