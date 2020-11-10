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
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	wt  = bufio.NewWriter(os.Stdout)
	mod = 1000000007
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

// example: ABC070D
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	g := NewGraph(n)
	for i := 0; i < n-1; i++ {
		a, b, c := getInt()-1, getInt()-1, getInt()
		g.AddEdge(a, b, c)
		g.AddEdge(b, a, c)
	}

	q, k := getInt(), getInt()-1
	g.DijkstraSearch(k)

	for i := 0; i < q; i++ {
		x, y := getInt()-1, getInt()-1
		puts(g.Cost[x] + g.Cost[y])
	}
}

type Edge struct {
	to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges    [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Pred     []int    // 遷移元
	Cost     []int    // 始点からの最小コスト
	RouteNum []int    // 始点から各頂点への最短経路数
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges:    make([][]Edge, n),
		Pred:     make([]int, n),
		Cost:     make([]int, n),
		RouteNum: make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
		g.Pred[i] = -1
		g.Cost[i] = 1 << 60
		g.RouteNum[i] = 0
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
func (g *Graph) DijkstraSearch(origin int) {
	// init
	g.Cost[origin] = 0
	g.RouteNum[origin] = 1
	done := make([]bool, len(g.Edges))
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
		done[from] = true
		for _, edge := range g.Edges[from] {
			to := edge.to
			if done[to] {
				continue
			}
			// 隣接頂点の最小コストを更新
			nextCost := v.priority + edge.cost
			if nextCost < g.Cost[to] {
				g.Cost[to] = nextCost
				g.Pred[to] = from
				g.RouteNum[to] = g.RouteNum[from]
				heap.Push(&pq, &Elem{
					value:    to,
					priority: nextCost,
				})
			} else if nextCost == g.Cost[to] {
				g.RouteNum[to] = (g.RouteNum[to] + g.RouteNum[from]) % mod
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
