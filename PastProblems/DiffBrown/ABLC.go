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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()

	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		a, b := getInt()-1, getInt()-1
		uf.Unite(a, b)
	}

	mp := map[int]bool{}
	for i := 0; i < n; i++ {
		mp[uf.Find(i)] = true
	}

	ans := len(mp) - 1
	puts(ans)
}
