package main

type UnionFind struct {
	par  []int
	rank []int
}

func (uf UnionFind) Init(n int) {
	for i := 0; i < n; i++ {
		uf.par[i] = i
		uf.rank[i] = 0
	}
}

func (uf UnionFind) Find(x int) int {
	for uf.par[x] != x {
		x, uf.par[x] = uf.par[x], uf.par[uf.par[x]]
	}
	return x
}

func (uf UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.rank[x] < uf.rank[y] {
		uf.par[x] = y
	} else {
		uf.par[y] = x
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func (uf UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// uf := UnionFind{make([]int, 3*n), make([]int, 3*n)}
// uf.Init();
