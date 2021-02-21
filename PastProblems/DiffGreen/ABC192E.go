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

func divCeil(a, b int) int {
	return (a + (b - 1)) / b
}

type Edge struct {
	to, cost, period int
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
		g.Cost[i] = inf
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost, period int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost, period})
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
		if g.Cost[from] < v.priority {
			continue
		}
		for _, edge := range g.Edges[from] {
			to := edge.to
			if done[to] {
				continue
			}
			// 隣接頂点の最小コストを更新
			wait := (edge.period - v.priority%edge.period) % edge.period
			nextCost := v.priority + wait + edge.cost
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

	n, m, x, y := getInt(), getInt(), getInt()-1, getInt()-1
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, t, k := getInt()-1, getInt()-1, getInt(), getInt()
		g.AddEdge(a, b, t, k)
		g.AddEdge(b, a, t, k)
	}

	g.DijkstraSearch(x)
	if g.Cost[y] == inf {
		puts(-1)
	} else {
		puts(g.Cost[y])
	}
}
