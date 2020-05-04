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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func grouping(friend [][]int, group []int, p, id int) {
	var dfs func(p int)
	dfs = func(p int) {
		for _, f := range friend[p] {
			if group[f] == 0 {
				group[f] = id
				dfs(f)
			}
		}
	}
	dfs(p)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, k := nextInt(), nextInt(), nextInt()
	friend, block := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		friend[i], block[i] = make([]int, 0), make([]int, 0)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		friend[a] = append(friend[a], b)
		friend[b] = append(friend[b], a)
	}
	for i := 0; i < k; i++ {
		c, d := nextInt()-1, nextInt()-1
		block[c] = append(block[c], d)
		block[d] = append(block[d], c)
	}

	// 友人の繋がりでグループに分類
	id := 0
	group := make([]int, n)
	cnt := map[int]int{}
	for i := 0; i < n; i++ {
		if group[i] == 0 {
			id++
			group[i] = id
			grouping(friend, group, i, id)
		}
		cnt[group[i]]++
	}

	// 同一グループに属し，互いにブロックしている相手の数
	sameGroupBlock := make([]int, n)
	for i := 0; i < n; i++ {
		for _, j := range block[i] {
			if group[i] == group[j] {
				sameGroupBlock[i]++
			}
		}
	}

	for i := 0; i < n; i++ {
		// 同一グループの自分以外の人数
		// - 直接の友人の数
		// - 同一グループに属していてブロックしている人数
		ans := cnt[group[i]] - 1 - len(friend[i]) - sameGroupBlock[i]
		if i < n-1 {
			putf("%d ", ans)
		} else {
			puts(ans)
		}
	}
}
