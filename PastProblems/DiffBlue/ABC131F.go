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
	max            = 100010
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

type Point struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	points := make([]Point, n)
	uf := NewUnionFind(2 * max)
	for i := 0; i < n; i++ {
		x, y := getInt(), getInt()
		points[i].x, points[i].y = x, y
		uf.Unite(x, y+max)
	}

	mx, my := map[int]int{}, map[int]int{}
	for i := 0; i < max; i++ {
		mx[uf.Find(i)]++
		my[uf.Find(i+max)]++
	}

	ans := 0
	for i := 0; i < 2*max; i++ {
		ans += mx[i] * my[i]
	}

	puts(ans - n)
}
