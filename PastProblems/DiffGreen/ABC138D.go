// 答え見た

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

var tree [][]int
var visited []bool
var cnt []int64

func dfs(now int, parentCnt int64) {
	cnt[now] += parentCnt
	visited[now] = true
	for _, v := range tree[now] {
		if !visited[v] {
			dfs(v, cnt[now])
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	q := nextInt()
	tree = make([][]int, n)
	for i := 0; i < n-1; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}
	cnt = make([]int64, n)
	for i := 0; i < q; i++ {
		p := nextInt()
		x := nextInt()
		cnt[p-1] += int64(x)
	}

	visited = make([]bool, n)
	dfs(0, 0)
	for _, v := range cnt {
		fmt.Println(v)
	}
}
