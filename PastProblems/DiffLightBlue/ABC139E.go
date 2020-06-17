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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func dfs(g *Graph, visited []bool, dp []int, cur int) int {
	if visited[cur] {
		return dp[cur]
	}
	visited[cur] = true
	cost := 1
	for _, e := range g.Edges[cur] {
		res := dfs(g, visited, dp, e.to)
		if res == -1 {
			return -1
		}
		cost = max(cost, res+1)
	}
	dp[cur] = cost
	return dp[cur]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n-1)
		for j := range a[i] {
			a[i][j] = nextInt() - 1
		}
	}

	// id[i][j]: 選手iと選手jの試合のid
	id := make([][]int, n)
	num := 0
	for i := 0; i < n; i++ {
		id[i] = make([]int, n)
		for j := range id[i] {
			if i < j {
				id[i][j] = num
				num++
			} else if i > j {
				id[i][j] = id[j][i]
			}
		}
	}
	// 試合を頂点とみなしたグラフ
	g := NewGraph(num)
	for i := 0; i < n; i++ {
		// 選手iの各選手との試合を，試合idで置き換える
		for j := 0; j < n-1; j++ {
			a[i][j] = id[i][a[i][j]]
		}
		// 試合を有向グラフに変換
		// 次の試合から今の試合への有向辺を張る
		for j := 0; j < n-2; j++ {
			g.AddEdge(a[i][j+1], a[i][j], 1)
		}
	}

	visited := make([]bool, num)
	// dp[i]: 頂点iからの最長経路長
	dp := make([]int, num)
	for i := range dp {
		dp[i] = -1
	}
	ans := 0
	for i := 0; i < num; i++ {
		res := dfs(g, visited, dp, i)
		if res == -1 {
			puts(-1)
			return
		}
		ans = max(ans, res)
	}

	puts(ans)
}
