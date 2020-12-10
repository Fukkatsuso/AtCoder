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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	a := getInts(3 * n)

	x := make([]int, 3*n)
	qx := new(PriorityQueueSmaller)
	for i, sum := 0, 0; i < 2*n; i++ {
		if qx.Len() < n {
			qx.push(a[i])
			sum += a[i]
		} else if qx.top() < a[i] {
			sum -= qx.pop()
			qx.push(a[i])
			sum += a[i]
		}
		x[i] = sum
	}

	y := make([]int, 3*n)
	qy := new(PriorityQueueBigger)
	for i, sum := 3*n-1, 0; i >= n; i-- {
		if qy.Len() < n {
			qy.push(a[i])
			sum += a[i]
		} else if qy.top() > a[i] {
			sum -= qy.pop()
			qy.push(a[i])
			sum += a[i]
		}
		y[i] = sum
	}

	ans := -(1 << 60)
	for i := n; i <= 2*n; i++ {
		ans = max(ans, x[i-1]-y[i])
	}

	puts(ans)
}

type PriorityQueueBigger []int

func (pq PriorityQueueBigger) Len() int {
	return len(pq)
}

// 要素を大きい順に構築する
func (pq PriorityQueueBigger) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueueBigger) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueBigger) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueueBigger) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// wrapper
func (pq *PriorityQueueBigger) push(x int) { heap.Push(pq, x) }
func (pq *PriorityQueueBigger) pop() int   { return heap.Pop(pq).(int) }
func (pq PriorityQueueBigger) top() int    { return pq[0] }

type PriorityQueueSmaller []int

func (pq PriorityQueueSmaller) Len() int {
	return len(pq)
}

// 要素を小さい順に構築する
func (pq PriorityQueueSmaller) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueueSmaller) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueSmaller) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueueSmaller) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// wrapper
func (pq *PriorityQueueSmaller) push(x int) { heap.Push(pq, x) }
func (pq *PriorityQueueSmaller) pop() int   { return heap.Pop(pq).(int) }
func (pq PriorityQueueSmaller) top() int    { return pq[0] }
