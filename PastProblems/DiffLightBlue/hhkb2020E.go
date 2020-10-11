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

// xの属する集合の要素数
func (uf *UnionFind) Size(x int) int {
	return -uf.Par[uf.Find(x)]
}

func index(i, j, w int) int {
	return i*w + j
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	free := 0
	for i := range s {
		s[i] = next()
		for j := range s[i] {
			if s[i][j] == '.' {
				free++
			}
		}
	}

	// マス(i,j)の縦方向のグループ, 横方向のグループ
	ufH, ufW := NewUnionFind(h*w), NewUnionFind(h*w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				continue
			}
			if i > 0 && s[i-1][j] == '.' {
				ufH.Unite(index(i-1, j, w), index(i, j, w))
			}
			if j > 0 && s[i][j-1] == '.' {
				ufW.Unite(index(i, j-1, w), index(i, j, w))
			}
		}
	}

	// mp[i]: 2^i(modつき)
	mp := make([]int, free+1)
	mp[0] = 1
	for i := 1; i <= free; i++ {
		mp[i] = mp[i-1] * 2
		mp[i] %= mod
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				continue
			}
			idx := index(i, j, w)
			// マス(i,j)に置いたときに照らせるマスの数
			size := ufH.Size(idx) + ufW.Size(idx) - 1
			// マス(i,j)が照らされるパターン数
			add := (mp[size] - 1 + mod) % mod
			// 上のパターン数は，他のグループのマスに証明を置くか否かを決めたときの組み合わせの数で，
			// 1つのパターンごとに他のグループの照らし方が2^(free-size)通りある
			add *= mp[free-size]
			add %= mod
			ans += add
			ans %= mod
		}
	}

	puts(ans)
}
