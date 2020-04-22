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

func isNotBridge(n int, table [][]bool, edge [2]int) bool {
	visited := make([]bool, n)
	// 0番目の頂点からスタート
	visited[0] = true

	// 辺を取り除く
	table[edge[0]][edge[1]], table[edge[1]][edge[0]] = false, false

	// bfs
	q := make([]int, 0)
	q = append(q, 0)
	for len(q) > 0 {
		from := q[0]
		for to := range table[from] {
			if table[from][to] && !visited[to] {
				q = append(q, to)
			}
		}
		visited[from] = true
		q = q[1:]
	}

	// 辺を復活
	table[edge[0]][edge[1]], table[edge[1]][edge[0]] = true, true

	// 全ての頂点を訪れたか?
	ok := true
	for i := range visited {
		ok = ok && visited[i]
	}
	return ok
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	edges := make([][2]int, m)
	table := make([][]bool, n)
	for i := range table {
		table[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		edges[i][0], edges[i][1] = a, b
		table[a][b], table[b][a] = true, true
	}

	ans := 0
	for i := range edges {
		if !isNotBridge(n, table, edges[i]) {
			ans++
		}
	}

	puts(ans)
}
