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

// レート，園児番号
type pqi struct {
	a, i int
}

type PriorityQueue1 []pqi

func (pq PriorityQueue1) Len() int            { return len(pq) }
func (pq PriorityQueue1) Less(i, j int) bool  { return pq[i].a > pq[j].a }
func (pq PriorityQueue1) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue1) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *PriorityQueue1) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// レート，園児番号，所属園
type pqj struct {
	a, i, d int
}

type PriorityQueue2 []pqj

// idxs[i]: 園iの最強園児のPriorityQueue2でのindex
var idxs [200000]int

func (pq PriorityQueue2) Len() int           { return len(pq) }
func (pq PriorityQueue2) Less(i, j int) bool { return pq[i].a < pq[j].a }
func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	idxs[pq[i].d], idxs[pq[j].d] = i, j
}
func (pq *PriorityQueue2) Push(x interface{}) {
	*pq = append(*pq, x.(pqj))
	idxs[x.(pqj).d] = len(*pq) - 1
}
func (pq *PriorityQueue2) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	idxs[x.d] = -1
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	// init idxs
	for i := range idxs {
		idxs[i] = -1
	}

	n, q := nextInt(), nextInt()
	rate := make([]int, n)

	// 園児の所属園
	mp := map[int]int{}
	// 各園で園児をまとめる
	infants := [200000]PriorityQueue1{}
	for i := range infants {
		infants[i] = PriorityQueue1{}
	}
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()-1
		rate[i] = a
		mp[i] = b
		heap.Push(&infants[b], pqi{a, i})
	}

	// 最強園児の集合
	maxrate := PriorityQueue2{}
	for i := range infants {
		if infants[i].Len() == 0 {
			continue
		}
		x := infants[i][0]
		heap.Push(&maxrate, pqj{x.a, x.i, i})
	}

	for i := 0; i < q; i++ {
		c, d := nextInt()-1, nextInt()-1

		// infants[from]から園児cを疑似的に削除・転園
		from := mp[c]
		mp[c] = d

		// maxrateの園fromを更新
		heap.Remove(&maxrate, idxs[from])
		idxs[from] = -1
		for infants[from].Len() > 0 {
			x := infants[from][0]
			// 本当に園fromにいる園児を扱う
			if mp[x.i] == from {
				heap.Push(&maxrate, pqj{x.a, x.i, from})
				break
			}
			heap.Pop(&infants[from])
		}

		// maxrateの園dを更新
		if idxs[d] >= 0 {
			heap.Remove(&maxrate, idxs[d])
			idxs[d] = -1
		}
		// infants[d]に追加
		heap.Push(&infants[d], pqi{rate[c], c})
		for {
			x := infants[d][0]
			if mp[x.i] == d {
				heap.Push(&maxrate, pqj{x.a, x.i, d})
				break
			}
			heap.Pop(&infants[d])
		}

		puts(maxrate[0].a)
	}
}
