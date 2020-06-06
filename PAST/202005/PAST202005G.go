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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const D = 1001
	const MAX = 500

	n, X, Y := nextInt(), nextInt()+MAX, nextInt()+MAX
	cannot := make([]bool, D*D+1)
	for i := 0; i < n; i++ {
		x, y := nextInt()+MAX, nextInt()+MAX
		cannot[x*D+y] = true
	}

	dx := []int{1, 0, -1, 1, -1, 0}
	dy := []int{1, 1, 1, 0, 0, -1}
	g := NewGraph(D*D + 1)
	for x := 0; x < D; x++ {
		for y := 0; y < D; y++ {
			if cannot[x*D+y] {
				continue
			}
			for k := 0; k < 6; k++ {
				nextX, nextY := x+dx[k], y+dy[k]
				if nextX < 0 || D <= nextX || nextY < 0 || D <= nextY {
					continue
				}
				if !cannot[nextX*D+nextY] {
					g.AddEdge(x*D+y, nextX*D+nextY, 1)
				}
			}
		}
	}

	g.DijkstraSearch(MAX*D + MAX)
	if ans := g.Cost[X*D+Y]; ans == 1<<60 {
		puts(-1)
	} else {
		puts(ans)
	}
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
