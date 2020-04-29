// dfsをミスっていた
// 他の人の提出を参考に修正してAC

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

func dfs(link [][]bool, visited []bool, prev, now int) bool {
	visited[now] = true
	for to := range link[now] {
		if to == prev || !link[now][to] {
			continue
		}
		if visited[to] {
			return false
		}
		if !dfs(link, visited, now, to) {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()

	link := make([][]bool, n)
	for i := range link {
		link[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		u, v := nextInt()-1, nextInt()-1
		link[u][v], link[v][u] = true, true
	}

	ans := 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] && dfs(link, visited, -1, i) {
			ans++
		}
	}

	puts(ans)
}
