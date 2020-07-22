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

func isChar(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	s1, s2 := next(), next()

	// 存在するか, 数字確定
	exist, bind := make([]bool, 26), make([]bool, 26)
	uf := NewUnionFind(26)
	for i := 0; i < n; i++ {
		id1, id2 := int(s1[i]-'A'), int(s2[i]-'A')
		if isChar(s1[i]) && isChar(s2[i]) {
			exist[id1], exist[id2] = true, true
			uf.Unite(id1, id2)
		} else if isChar(s1[i]) {
			exist[id1] = true
			bind[id1] = true
		} else if isChar(s2[i]) {
			exist[id2] = true
			bind[id2] = true
		}
	}

	ans := 1
	checked := make([]bool, 26)
	for i := 0; i < 26; i++ {
		if !exist[i] || checked[uf.Find(i)] {
			continue
		}
		binded, hasTop := false, false
		for j := 0; j < 26; j++ {
			binded = binded || (uf.Same(i, j) && bind[j])
			hasTop = hasTop || (uf.Same(i, j) && (int(s1[0]-'A') == j || int(s2[0]-'A') == j))
		}
		if !binded && hasTop {
			ans *= 9
		} else if !binded {
			ans *= 10
		}
		checked[uf.Find(i)] = true
	}

	puts(ans)
}
