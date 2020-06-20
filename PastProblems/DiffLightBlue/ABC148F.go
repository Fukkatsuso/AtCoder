// 解説AC

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
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges: make([][]Edge, n),
		Cost:  make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
		g.Cost[i] = 1 << 60
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost})
}

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

	n, u, v := nextInt(), nextInt()-1, nextInt()-1
	g1 := NewGraph(n) // 逃げる（高橋）
	g2 := NewGraph(n) // 追う（青木）
	for i := 0; i < n-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		g1.AddEdge(a, b, 1)
		g1.AddEdge(b, a, 1)
		g2.AddEdge(a, b, 1)
		g2.AddEdge(b, a, 1)
	}

	if len(g1.Edges[u]) == 1 && g1.Edges[u][0].to == v {
		puts(0)
		return
	}

	g1.DijkstraSearch(u)
	g2.DijkstraSearch(v)

	// 青木から一番遠い葉を探す
	x := v
	for i := 0; i < n; i++ {
		if g1.Cost[i] < g2.Cost[i] && g2.Cost[x] < g2.Cost[i] {
			x = i
		}
	}
	// 青木が最後に到着する点
	last := g1.Edges[x][0].to
	puts(g2.Cost[last])
}
