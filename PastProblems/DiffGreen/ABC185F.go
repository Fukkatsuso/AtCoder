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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, q := getInt(), getInt()
	seg := NewSegTree(n, 0, func(a, b SegType) SegType { return a ^ b })
	for i := 0; i < n; i++ {
		a := getInt()
		seg.Set(i, a)
	}
	seg.Build()

	for ; q > 0; q-- {
		t, x, y := getInt(), getInt(), getInt()
		if t == 1 {
			a := seg.Get(x - 1)
			seg.Update(x-1, a^y)
		} else {
			ans := seg.Eval(x-1, y)
			puts(ans)
		}
	}
}
