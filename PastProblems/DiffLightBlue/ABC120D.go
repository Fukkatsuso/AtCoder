// Union-Findの要素数計算実装に失敗してWA
// 修正してAC

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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	a, b := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = nextInt()-1, nextInt()-1
	}

	ans := make([]int, m+1)
	ans[m] = n * (n - 1) / 2
	uf := NewUnionFind(n)
	for i := m - 1; i >= 0; i-- {
		if uf.Same(a[i], b[i]) {
			ans[i] = ans[i+1]
		} else {
			ans[i] = ans[i+1] - uf.Size(a[i])*uf.Size(b[i])
			uf.Unite(a[i], b[i])
		}
	}

	for i := 1; i <= m; i++ {
		puts(ans[i])
	}
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

func (uf *UnionFind) Size(x int) int {
	return -uf.Par[uf.Find(x)]
}
