package main

import (
	"container/heap"
)

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 要素を大きい順に構築する
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// wrapper
func (pq *PriorityQueue) push(x int) { heap.Push(pq, x) }
func (pq *PriorityQueue) pop() int   { return heap.Pop(pq).(int) }
func (pq PriorityQueue) top() int    { return pq[0] }
