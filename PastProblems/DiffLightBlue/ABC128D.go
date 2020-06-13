// AC
// 余裕あるのに計算量を意識しすぎて複雑実装になってしまった

package main

import (
	"bufio"
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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
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

type Heap struct {
	size  int
	elems [100]int
}

func (heap *Heap) Size() int {
	return heap.size
}

func (heap *Heap) Push(x int) {
	i := heap.size

	for i > 0 {
		p := (i - 1) / 2
		if heap.elems[p] <= x {
			break
		}
		heap.elems[i] = heap.elems[p]
		i = p
	}

	heap.elems[i] = x
	heap.size++
}

func (heap *Heap) Pop() int {
	// 最小値
	ret := heap.elems[0]
	heap.size--
	x := heap.elems[heap.size]

	i := 0
	for i*2+1 < heap.size {
		a, b := i*2+1, i*2+2
		if b < heap.size && heap.elems[b] < heap.elems[a] {
			a = b
		}
		if heap.elems[a] >= x {
			break
		}

		heap.elems[i] = heap.elems[a]
		i = a
	}

	heap.elems[i] = x

	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	v := nextInts(n)

	ans := 0
	suml, minl := 0, new(Heap)
	for l := -1; l < n; l++ {
		if l >= 0 {
			suml += v[l]
			minl.Push(v[l])
		}
		sumr, minr := 0, new(Heap)
		for r := n; l < r && (l+1)+(n-r) <= k; r-- {
			if r < n {
				sumr += v[r]
				minr.Push(v[r])
			}
			val := suml + sumr
			// 取り出す要素
			remr, reml := make([]int, 0), make([]int, 0)
			for d := k - (l + 1) - (n - r); d > 0; d-- {
				if minl.Size() > 0 && minl.elems[0] < 0 && minr.Size() > 0 && minr.elems[0] < 0 {
					if minl.elems[0] < minr.elems[0] {
						rem := minl.Pop()
						reml = append(reml, rem)
						val -= rem
					} else {
						rem := minr.Pop()
						remr = append(remr, rem)
						val -= rem
					}
				} else if minl.Size() > 0 && minl.elems[0] < 0 {
					rem := minl.Pop()
					reml = append(reml, rem)
					val -= rem
				} else if minr.Size() > 0 && minr.elems[0] < 0 {
					rem := minr.Pop()
					remr = append(remr, rem)
					val -= rem
				}
			}
			ans = max(ans, val)
			// 復元
			for i := range reml {
				minl.Push(reml[i])
			}
			for i := range remr {
				minr.Push(remr[i])
			}
		}
	}

	puts(ans)
}
