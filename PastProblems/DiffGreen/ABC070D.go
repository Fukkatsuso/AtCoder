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

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

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
	pq[i].index = i
	pq[j].index = j
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	edges := make([][]Edge, n)
	for i := range edges {
		edges[i] = make([]Edge, 0)
	}
	for i := 0; i < n-1; i++ {
		a, b, c := nextInt3()
		a--
		b--
		edges[a] = append(edges[a], Edge{b, c})
		edges[b] = append(edges[b], Edge{a, c})
	}

	q, k := nextInt2()
	k--

	// 頂点kから各点への最短距離を求める
	dist := make([]int, n)
	for i := range dist {
		if i != k {
			dist[i] = 1 << 60 // INF
		}
	}
	// ダイクストラ法
	done := make([]bool, n)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	// 始点k
	heap.Push(&pq, &Elem{
		value:    k,
		priority: 0,
	})
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Elem)
		done[v.value.(int)] = true
		for _, next := range edges[v.value.(int)] {
			if done[next.to] {
				continue
			}
			cost := v.priority + next.cost
			if cost < dist[next.to] {
				dist[next.to] = cost
				heap.Push(&pq, &Elem{
					value:    next.to,
					priority: cost,
				})
			}
		}
	}

	for i := 0; i < q; i++ {
		x, y := nextInt2()
		x--
		y--
		puts(dist[x] + dist[y])
	}
}
