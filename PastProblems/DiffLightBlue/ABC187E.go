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
	from, to int
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
	g.Edges[from] = append(g.Edges[from], Edge{from, to})
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	g := NewGraph(n)
	a, b := make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i], b[i] = getInt()-1, getInt()-1
		g.AddEdge(a[i], b[i])
		g.AddEdge(b[i], a[i])
	}

	toChild := map[Edge]bool{}
	que := []int{}
	que = append(que, 0)
	for visited := make([]bool, n); len(que) > 0; {
		cur := que[0]
		que = que[1:]
		visited[cur] = true
		for _, e := range g.Edges[cur] {
			if !visited[e.to] {
				toChild[e] = true
				que = append(que, e.to)
			}
		}
	}

	dp := make([]int, n)
	q := getInt()
	for i := 0; i < q; i++ {
		t, e, x := getInt(), getInt()-1, getInt()
		from, not := a[e], b[e]
		if t == 2 {
			from, not = not, from
		}
		if toChild[Edge{from, not}] {
			// notの子孫以外に+x
			dp[0] += x
			dp[not] -= x
		} else {
			// fromの子孫に+x
			dp[from] += x
		}
	}

	// rootから累積和
	que = append(que, 0)
	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		for _, e := range g.Edges[cur] {
			if toChild[e] {
				dp[e.to] += dp[cur]
				que = append(que, e.to)
			}
		}
	}

	for i := 0; i < n; i++ {
		puts(dp[i])
	}
}
