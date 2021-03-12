// 解説AC
// 誤読（条件の誤解）でWA，考察はOK

package main

import (
	"bufio"
	"container/heap"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Edge struct {
	from, to, cost, id int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Pred  []int    // 遷移元とをつなぐ辺のID
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
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost, id int) {
	g.Edges[from] = append(g.Edges[from], Edge{from, to, cost, id})
}

// value: 頂点の番号
// priority: 始点からのコスト
type Elem struct {
	value    interface{}
	priority int
	index    int
}

type PriorityQueue []*Elem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	elem := x.(*Elem)
	elem.index = n
	*pq = append(*pq, elem)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	elem := old[n-1]
	old[n-1] = nil
	elem.index = -1
	*pq = old[0 : n-1]
	return elem
}

// 0-indexed
// eID: 取り除く辺のID
func (g *Graph) DijkstraSearch(origin, eID int) {
	// init
	for i := 0; i < len(g.Edges); i++ {
		g.Pred[i] = -1
		g.Cost[i] = inf
	}
	g.Cost[origin] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Elem{
		value:    origin,
		priority: 0,
	})

	// seaech
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Elem)
		from := v.value.(int)
		if g.Cost[from] < v.priority {
			continue
		}
		for _, edge := range g.Edges[from] {
			if edge.id == eID {
				continue
			}
			to := edge.to
			// 隣接頂点の最小コストを更新
			nextCost := v.priority + edge.cost
			if nextCost < g.Cost[to] {
				g.Cost[to] = nextCost
				g.Pred[to] = edge.from
				heap.Push(&pq, &Elem{
					value:    to,
					priority: nextCost,
				})
			}
		}
	}
}

// toへ向かう最短パスを返す
func (g *Graph) MinPath(to int) []int {
	path := make([]int, 0)
	// 終点から探索
	path = append(path, to)
	for node := to; g.Pred[node] != -1; {
		node = g.Pred[node]
		path = append(path, node)
	}
	// 配列を反転
	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}
	return path
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	g := NewGraph(n)
	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		a, b := getInt()-1, getInt()-1
		g.AddEdge(a, b, 1, i)
		edges[i] = Edge{a, b, 1, i}
	}

	// i番目の辺を含む最短閉路を探す
	// あれば出力して終了
	for i := 0; i < m; i++ {
		g.DijkstraSearch(edges[i].to, i)
		if g.Cost[edges[i].from] == inf {
			continue
		}
		path := g.MinPath(edges[i].from)
		// 辺eの両端点がともにpath上の点集合vに含まれていれば，端点の次数を加算
		// 次数が3以上になればNG
		v := make([]bool, n)
		for _, p := range path {
			v[p] = true
		}
		degree := make([]int, n)
		ok := true
		for _, e := range edges {
			if v[e.from] && v[e.to] {
				degree[e.from]++
				degree[e.to]++
			}
		}
		for _, d := range degree {
			ok = ok && d <= 2
		}
		if !ok {
			continue
		}
		// output
		puts(len(path))
		for _, pid := range path {
			puts(pid + 1)
		}
		return
	}

	puts(-1)
}
