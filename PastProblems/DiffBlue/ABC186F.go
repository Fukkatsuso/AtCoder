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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
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

type Point struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, m := getInt(), getInt(), getInt()
	// ymin[i]: i行目の障害物で一番小さいy
	// xmin[i]: i列目の障害物で一番小さいx
	ymin, xmin := make([]int, h), make([]int, w)
	for i := range ymin {
		ymin[i] = w
	}
	for i := range xmin {
		xmin[i] = h
	}
	p := make([]Point, m)
	for i := 0; i < m; i++ {
		x, y := getInt()-1, getInt()-1
		p[i].x, p[i].y = x, y
		ymin[x] = min(ymin[x], y)
		xmin[y] = min(xmin[y], x)
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})

	ans := 0
	for i := range xmin {
		ans += xmin[i]
		if xmin[i] == 0 {
			break
		}
	}

	seg := NewSegTree(w+1, 0, func(a, b SegType) SegType { return a + b })
	for i := ymin[0]; i < w; i++ {
		seg.Set(i, 1)
	}
	seg.Build()
	for i, pIdx := 1, 0; i < h; i++ {
		add := seg.Eval(0, ymin[i])
		for pIdx < m && p[pIdx].x <= i {
			if seg.Get(p[pIdx].y) == 0 {
				seg.Update(p[pIdx].y, 1)
			}
			pIdx++
		}
		if seg.Get(0) == 0 {
			ans += add
		}
	}

	puts(ans)
}
