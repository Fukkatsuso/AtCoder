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

type Job struct {
	a, b int
}

type Jobs []Job

func (a Jobs) Len() int           { return len(a) }
func (a Jobs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Jobs) Less(i, j int) bool { return a[i].a < a[j].a }

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	jobs := make(Jobs, n)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		jobs[i] = Job{a, b}
	}
	sort.Sort(jobs)

	ans := 0
	pq := make(PriorityQueue, 0)
	for day, id := m-1, 0; day >= 0; day-- {
		for id < n && day+jobs[id].a <= m {
			heap.Push(&pq, jobs[id].b)
			id++
		}
		if pq.Len() > 0 {
			ans += heap.Pop(&pq).(int)
		}
	}

	puts(ans)
}
