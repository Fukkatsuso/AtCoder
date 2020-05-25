package main

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
