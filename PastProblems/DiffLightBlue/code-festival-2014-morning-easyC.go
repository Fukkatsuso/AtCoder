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
	inf            = 1 << 50
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
	Cost  []int    // 始点からの最小コスト
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int, edges [][]Edge) *Graph {
	g := &Graph{
		Edges: edges,
		Cost:  make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Cost[i] = inf
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost})
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
				heap.Push(&pq, &Elem{
					value:    to,
					priority: nextCost,
				})
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	s, t := nextInt()-1, nextInt()-1

	edges := make([][]Edge, n)
	for i := range edges {
		edges[i] = make([]Edge, 0)
	}
	for i := 0; i < m; i++ {
		x, y, d := nextInt()-1, nextInt()-1, nextInt()
		edges[x] = append(edges[x], Edge{y, d})
		edges[y] = append(edges[y], Edge{x, d})
	}

	gs, gt := NewGraph(n, edges), NewGraph(n, edges)
	gs.DijkstraSearch(s)
	gt.DijkstraSearch(t)

	for i := 0; i < n; i++ {
		if gs.Cost[i] != inf && gs.Cost[i] == gt.Cost[i] {
			puts(i + 1)
			return
		}
	}

	puts(-1)
}
