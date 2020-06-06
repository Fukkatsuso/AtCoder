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

type UnionFind struct {
	Par  []int
	Rank []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		Par:  make([]int, n),
		Rank: make([]int, n),
	}
	uf.Init(n)
	return uf
}

func (uf *UnionFind) Init(n int) {
	for i := 0; i < n; i++ {
		uf.Par[i] = i
		uf.Rank[i] = 0
	}
}

func (uf *UnionFind) Find(x int) int {
	for uf.Par[x] != x {
		x, uf.Par[x] = uf.Par[x], uf.Par[uf.Par[x]]
	}
	return x
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.Rank[x] < uf.Rank[y] {
		uf.Par[x] = y
	} else {
		uf.Par[y] = x
		if uf.Rank[x] == uf.Rank[y] {
			uf.Rank[x]++
		}
	}
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	p := make([]int, n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt() - 1
		x[p[i]] = i
	}

	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		x, y := nextInt()-1, nextInt()-1
		uf.Unite(x, y)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if uf.Same(i, x[i]) {
			ans++
		}
	}

	puts(ans)
}
