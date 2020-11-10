// 解説AC

package main

import (
	"bufio"
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

// 求めたい値の型
type SegType = int

type SegTree struct {
	size  int
	data  []SegType
	f     func(a, b SegType) SegType
	unity SegType
}

func NewSegTree(n int, unity SegType, f func(a, b SegType) SegType) *SegTree {
	size := 1
	for size < n {
		size *= 2
	}
	data := make([]SegType, size*2)
	for i := range data {
		data[i] = unity
	}
	seg := SegTree{
		size:  size,
		data:  data,
		f:     f,
		unity: unity,
	}
	return &seg
}

func (seg *SegTree) Set(idx int, v SegType) {
	seg.data[idx+seg.size] = v
}

func (seg *SegTree) Build() {
	for i := seg.size - 1; i > 0; i-- {
		seg.data[i] = seg.f(seg.data[i*2], seg.data[i*2+1])
	}
}

func (seg *SegTree) Update(idx int, v SegType) {
	seg.Set(idx, v)
	for i := (idx + seg.size) / 2; i > 0; i >>= 1 {
		seg.data[i] = seg.f(seg.data[i*2], seg.data[i*2+1])
	}
}

// [l, r)
func (seg *SegTree) Eval(l, r int) SegType {
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

func (seg *SegTree) Get(idx int) SegType {
	return seg.data[idx+seg.size]
}

type Box struct {
	w, h int
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

	n := nextInt()
	b := make([]Box, n)
	for i := range b {
		w, h := nextInt(), nextInt()
		b[i] = Box{w, h}
	}

	// hの昇順, wの降順
	sort.Slice(b, func(i, j int) bool {
		if b[i].h == b[j].h {
			return b[i].w > b[j].w
		}
		return b[i].h < b[j].h
	})

	// dp配列をセグ木にのせる
	seg := NewSegTree(100010, 0, func(a int, b int) int {
		return max(a, b)
	})
	for i := 0; i < n; i++ {
		val := seg.Eval(1, b[i].w) + 1
		seg.Update(b[i].w, val)
	}

	ans := seg.Eval(1, 100010)
	puts(ans)
}
