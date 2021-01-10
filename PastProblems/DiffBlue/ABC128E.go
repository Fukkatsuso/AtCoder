// 解説AC

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
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

type Event struct {
	t, op, x int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, q := getInt(), getInt()
	events := make([]Event, 2*n)
	for i := 0; i < n; i++ {
		s, t, x := getInt(), getInt(), getInt()
		events[2*i] = Event{s - x, 1, x}
		events[2*i+1] = Event{t - x, -1, x}
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].t < events[j].t
	})

	d := getInts(q)

	pq := PriorityQueue{}
	mp := map[int]int{}
	eventId := 0
	for i := 0; i < q; i++ {
		// i番目の人に影響のある範囲でイベントを追加し，削除するイベントのxはmpに記録
		for eventId < 2*n && events[eventId].t <= d[i] {
			e := events[eventId]
			if e.op == 1 {
				pq.push(e.x)
			} else {
				mp[e.x]++
			}
			eventId++
		}
		// イベント削除(キューの先頭にあるのが削除するイベントだったら邪魔)
		for pq.Len() > 0 && mp[pq.top()] > 0 {
			mp[pq.pop()]--
		}

		// 残っているイベントのうち最小のxが答え
		if pq.Len() == 0 {
			puts(-1)
		} else {
			puts(pq.top())
		}
	}
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
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
