package main

import "fmt"

type Heap struct {
	size  int
	elems [100000]int64
}

type MyHeap interface {
	Size() int
	Push(int64)
	Pop() int64
}

func (heap *Heap) Size() int {
	return heap.size
}

func (heap *Heap) Push(x int64) {
	i := heap.size

	for i > 0 {
		p := (i - 1) / 2

		if heap.elems[p] >= x {
			break
		}

		heap.elems[i] = heap.elems[p]
		i = p
	}

	heap.elems[i] = x
	heap.size++
}

func (heap *Heap) Pop() int64 {
	ret := heap.elems[0]
	heap.size--
	x := heap.elems[heap.size]

	i := 0
	for i*2+1 < heap.size {
		a, b := i*2+1, i*2+2
		if b < heap.size && heap.elems[b] > heap.elems[a] {
			a = b
		}
		if heap.elems[a] <= x {
			break
		}

		heap.elems[i] = heap.elems[a]
		i = a
	}

	heap.elems[i] = x

	return ret
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	pq := new(Heap)
	for i := 0; i < n; i++ {
		var a int64
		fmt.Scan(&a)
		pq.Push(a)
	}

	for i := 0; i < m; i++ {
		max := pq.Pop()
		pq.Push(max / 2)
	}

	fmt.Println(sum(pq.elems))
}

func sum(a [100000]int64) int64 {
	ret := int64(0)
	for _, v := range a {
		ret += v
	}
	return ret
}
