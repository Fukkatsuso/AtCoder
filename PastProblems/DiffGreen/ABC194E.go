// 解説AC
// 参考: https://marco-note.net/abc-194-work-log/

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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	a := getInts(n)

	seg := NewSegTree(2000000, inf, func(a, b SegType) SegType {
		return min(a, b)
	})
	cnt := map[int]int{}
	for i := 0; i < m-1; i++ {
		add := a[i]
		seg.Update(add, inf)
		cnt[add]++
	}
	seg.Build()

	ans := inf
	for i := m - 1; i < n; i++ {
		// 部分列の末尾を追加
		add := a[i]
		seg.Update(add, inf)
		cnt[add]++

		ans = min(ans, seg.Query(0, n+1))

		// 部分列の先端を削除
		del := a[i-m+1]
		if cnt[del] == 1 {
			seg.Update(del, del)
		}
		cnt[del]--
	}

	puts(ans)
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
		data[i] = i - size
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
