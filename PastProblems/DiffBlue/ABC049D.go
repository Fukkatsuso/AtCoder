// 解説AC

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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type UnionFind struct {
	Par []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		Par: make([]int, n),
	}
	uf.Init(n)
	return uf
}

func (uf *UnionFind) Init(n int) {
	for i := 0; i < n; i++ {
		uf.Par[i] = -1
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Par[x] < 0 {
		return x
	}
	uf.Par[x] = uf.Find(uf.Par[x])
	return uf.Par[x]
}

func (uf *UnionFind) Unite(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x == y {
		return
	}

	if uf.Par[x] > uf.Par[y] {
		x, y = y, x
	}
	uf.Par[x] += uf.Par[y]
	uf.Par[y] = x
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// xの属する集合の要素数
func (uf *UnionFind) Size(x int) int {
	return -uf.Par[uf.Find(x)]
}

type Groups struct {
	a, b int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k, l := getInt(), getInt(), getInt()
	// 道路
	ufa := NewUnionFind(n)
	for i := 0; i < k; i++ {
		p, q := getInt()-1, getInt()-1
		ufa.Unite(p, q)
	}
	// 鉄道
	ufb := NewUnionFind(n)
	for i := 0; i < l; i++ {
		r, s := getInt()-1, getInt()-1
		ufb.Unite(r, s)
	}

	mp := map[Groups]int{}
	for i := 0; i < n; i++ {
		g := Groups{
			a: ufa.Find(i),
			b: ufb.Find(i),
		}
		mp[g]++
	}

	for i := 0; i < n; i++ {
		g := Groups{
			a: ufa.Find(i),
			b: ufb.Find(i),
		}
		if i < n-1 {
			putf("%d ", mp[g])
		} else {
			puts(mp[g])
		}
	}
}
