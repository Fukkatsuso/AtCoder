// 解説AC
// 参考
// https://drken1215.hatenablog.com/entry/2020/09/27/080300

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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
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

type SegTree struct {
	n     int
	size  int
	unity int
	f     func(int, int) int
	data  []int
}

func NewSegTree(n int, unity int, f func(int, int) int) *SegTree {
	size := 1
	for size < n {
		size <<= 1
	}
	data := make([]int, size*2)
	for i := range data {
		data[i] = unity
	}
	seg := SegTree{
		n:     n,
		size:  size,
		unity: unity,
		f:     f,
		data:  data,
	}
	return &seg
}

// f([l, r)), 0-indexed
func (seg *SegTree) Get(l, r int) int {
	vl, vr := seg.unity, seg.unity
	for l, r = l+seg.size, r+seg.size; l < r; l, r = l>>1, r>>1 {
		if l&1 > 0 {
			vl = seg.f(vl, seg.data[l])
			l++
		}
		if r&1 > 0 {
			r--
			vr = seg.f(seg.data[r], vr)
		}
	}
	return seg.f(vl, vr)
}

// update a, 0-indexed
func (seg *SegTree) Update(a, x int) {
	k := a + seg.size
	seg.data[k] = x
	for k >>= 1; k > 0; k >>= 1 {
		seg.data[k] = seg.f(seg.data[k*2], seg.data[k*2+1])
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const MAX = 510000

	n, k := nextInt(), nextInt()
	a := nextInts(n)

	seg := NewSegTree(MAX, 0, func(x, y int) int { return max(x, y) })
	for i := 0; i < n; i++ {
		l, r := max(0, a[i]-k), min(MAX, a[i]+k+1)
		val := seg.Get(l, r)
		seg.Update(a[i], val+1)
	}

	ans := seg.Get(0, MAX)
	puts(ans)
}
