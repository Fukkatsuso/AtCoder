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
	mod            = 1000000007
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

func P(n, k int) int {
	res := 1
	for ; k >= 1; k-- {
		res *= n
		res %= mod
		n--
	}
	return res
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

func dfs(g *Graph, k, cur, prev int, ans *int) {
	if prev >= 0 {
		*ans *= P(max(0, k-2), len(g.Edges[cur])-1)
	} else {
		*ans *= P(k-1, len(g.Edges[cur]))
	}
	*ans %= mod
	for _, e := range g.Edges[cur] {
		if e.to != prev {
			dfs(g, k, e.to, cur, ans)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	g := NewGraph(n)
	for i := 0; i < n-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		g.AddEdge(a, b, 1)
		g.AddEdge(b, a, 1)
	}

	ans := k
	dfs(g, k, 0, -1, &ans)

	puts(ans)
}
