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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt2()
	d := NewDijkstra(n)
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		d.AddEdge(a, b, 1)
		d.AddEdge(b, a, 1)
	}
	d.Search(0)

	for i := 1; i < n; i++ {
		if d.Pred[i] < 0 {
			puts("No")
			return
		}
	}

	puts("Yes")
	for i := 1; i < n; i++ {
		puts(d.Pred[i] + 1)
	}
}

// 辺
type Edge struct {
	to, cost int
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

func newSlice(n, x int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = x
	}
	return slice
}

type Dijkstra struct {
	NodeNum int      // 頂点数
	Origin  int      // 始点
	InfCost int      // 無限大とみなすコストの値
	Pred    []int    // 遷移元
	Cost    []int    // 始点からの最小コスト
	Edges   [][]Edge // 頂点の隣接リスト
}

func NewDijkstra(n int) *Dijkstra {
	edges := make([][]Edge, n)
	for i := range edges {
		edges[i] = make([]Edge, 0)
	}
	inf := 1 << 60
	d := &Dijkstra{
		NodeNum: n,
		InfCost: inf,
		Pred:    newSlice(n, -1),
		Cost:    newSlice(n, inf),
		Edges:   edges,
	}
	return d
}

// fromからtoへの重み付きの辺を追加
func (d *Dijkstra) AddEdge(from, to, cost int) {
	d.Edges[from] = append(d.Edges[from], Edge{to, cost})
}

// 0-indexed
func (d *Dijkstra) Search(origin int) {
	// init
	done := make([]bool, d.NodeNum)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Elem{
		value:    origin,
		priority: 0,
	})

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Elem)
		now := v.value.(int)
		done[now] = true
		for _, edge := range d.Edges[now] {
			if done[edge.to] {
				continue
			}
			// 隣接頂点の最小コストを更新
			nextCost := v.priority + edge.cost
			if nextCost < d.Cost[edge.to] {
				d.Cost[edge.to] = nextCost
				d.Pred[edge.to] = now
				heap.Push(&pq, &Elem{
					value:    edge.to,
					priority: nextCost,
				})
			}
		}
	}
}
