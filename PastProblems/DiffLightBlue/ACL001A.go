// 解説AC（ほぼ写経）
// 参考
// https://drken1215.hatenablog.com/entry/2020/09/23/080100

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
	x, y, id int
}

type Points []Point

func (a Points) Len() int      { return len(a) }
func (a Points) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Points) Less(i, j int) bool {
	return a[i].x < a[j].x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	p := make(Points, n)
	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		p[i] = Point{x, y, i}
	}
	sort.Sort(p)

	uf := NewUnionFind(n)
	// {y, id}のスタック
	st := make([][2]int, 0)
	ymin := [2]int{1 << 30, -1}
	for i := 0; i < n; i++ {
		id := p[i].id
		if ymin[0] > p[i].y {
			st = append(st, [2]int{ymin[0], ymin[1]})
			ymin[0], ymin[1] = p[i].y, id
		} else {
			uf.Unite(id, ymin[1])
			for l := len(st); l > 0 && st[l-1][0] < p[i].y; l-- {
				uf.Unite(id, st[l-1][1])
				st = st[:l-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		puts(uf.Size(i))
	}
}
