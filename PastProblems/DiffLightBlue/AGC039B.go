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

// 二部グラフ判定
func isBipartite(s []string) bool {
	n := len(s)
	// -1:白, 0:無色, 1:黒
	colors := make([]int, n)
	var dfs func(i, color int) bool
	dfs = func(i, color int) bool {
		colors[i] = color
		for j := range s[i] {
			if s[i][j] == '0' {
				continue
			}
			if colors[j] == color {
				return false
			}
			if colors[j] == 0 && !dfs(j, -color) {
				return false
			}
		}
		return true
	}
	for i := range s {
		if colors[i] == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1 << 50

	n := nextInt()
	s := make([]string, n)
	for i := range s {
		s[i] = next()
	}

	if !isBipartite(s) {
		puts(-1)
		return
	}

	// ワーシャルフロイド
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		dist[i][i] = 0
		for j := 0; j < n; j++ {
			if s[i][j] == '1' {
				dist[i][j] = 1
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	// 直径
	d := 0
	for i := range dist {
		d = max(d, max(dist[i]...))
	}
	puts(d + 1)
}
