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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		edge[a][b], edge[b][a] = true, true
	}

	ans := 0

	var dfs func(visited []bool, cur int)
	dfs = func(visited []bool, cur int) {
		visited[cur] = true
		finish := true
		for i := 0; i < n; i++ {
			if i == cur {
				continue
			}
			finish = finish && visited[i]
			if edge[cur][i] && !visited[i] {
				dfs(visited, i)
			}
		}
		if finish {
			ans++
		}
		visited[cur] = false
	}
	dfs(make([]bool, n), 0)

	puts(ans)
}
