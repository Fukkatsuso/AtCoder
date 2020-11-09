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
	mod            = 998244353
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

func fact(n int) int {
	res := 1
	for ; n > 0; n-- {
		res *= n
		res %= mod
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	a := make([][]int, n)
	for i := range a {
		a[i] = nextInts(n)
	}

	// 行，列のグループ
	col, row := NewUnionFind(n), NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// col
			ok := true
			for x := 0; x < n; x++ {
				ok = ok && (a[i][x]+a[j][x] <= k)
			}
			if ok {
				col.Unite(i, j)
			}
			// row
			ok = true
			for x := 0; x < n; x++ {
				ok = ok && (a[x][i]+a[x][j] <= k)
			}
			if ok {
				row.Unite(i, j)
			}
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		if col.Find(i) == i {
			ans *= fact(col.Size(i))
			ans %= mod
		}
	}
	for i := 0; i < n; i++ {
		if row.Find(i) == i {
			ans *= fact(row.Size(i))
			ans %= mod
		}
	}
	puts(ans)
}
