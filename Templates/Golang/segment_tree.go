package main

import "fmt"

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
func (seg *SegTree) Query(l, r int) SegType {
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
	n := 10
	seg := NewSegTree(n, 0, func(a, b SegType) SegType {
		return a + b
	})
	for i := 0; i < n; i++ {
		seg.Set(i, i)
	}
	seg.Build()

	// 区間[0,5)の和=10
	fmt.Println(0, 5, seg.Query(0, 5))

	// [10,1,2,3,4,...]に更新, 和=20
	seg.Update(0, 10)
	fmt.Println(0, 5, seg.Query(0, 5))
}
