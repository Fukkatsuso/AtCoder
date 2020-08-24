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

	inf = 1 << 50
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

func connect(s []string, id [][]int, h, w int) int {
	di := []int{0, 1, 0, -1}
	dj := []int{1, 0, -1, 0}
	bfs := func(si, sj, x int) {
		que := make([][2]int, 0)
		id[si][sj] = x
		que = append(que, [2]int{si, sj})
		for len(que) > 0 {
			i, j := que[0][0], que[0][1]
			que = que[1:]
			for k := 0; k < 4; k++ {
				nextI, nextJ := i+di[k], j+dj[k]
				if nextI < 0 || h <= nextI || nextJ < 0 || w <= nextJ {
					continue
				}
				if s[nextI][nextJ] == '#' || id[nextI][nextJ] > 0 {
					continue
				}
				id[nextI][nextJ] = x
				que = append(que, [2]int{nextI, nextJ})
			}
		}
	}
	x := 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if id[i][j] > 0 || s[i][j] == '#' {
				continue
			}
			bfs(i, j, x)
			x++
		}
	}
	return x - 1
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := nextInt(), nextInt()
	ch, cw := nextInt()-1, nextInt()-1
	dh, dw := nextInt()-1, nextInt()-1
	s := make([]string, h)
	for i := range s {
		s[i] = next()
	}

	// 連結している道を同じidにする
	id := make([][]int, h)
	for i := range id {
		id[i] = make([]int, w)
	}
	maxID := connect(s, id, h, w)

	g := NewGraph(maxID + 1)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				continue
			}
			// 5x5を探索
			for dI := -2; dI <= 2; dI++ {
				for dJ := -2; dJ <= 2; dJ++ {
					nextI, nextJ := i+dI, j+dJ
					if nextI < 0 || h <= nextI || nextJ < 0 || w <= nextJ {
						continue
					}
					if s[nextI][nextJ] == '#' || id[i][j] == id[nextI][nextJ] {
						continue
					}
					g.AddEdge(id[i][j], id[nextI][nextJ], 1)
				}
			}
		}
	}
	g.DijkstraSearch(id[ch][cw])

	if g.Cost[id[dh][dw]] == inf {
		puts(-1)
	} else {
		puts(g.Cost[id[dh][dw]])
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
