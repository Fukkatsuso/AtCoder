// AC
// 解説の楽な書き方に変更した

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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1 << 60

	n, m := nextInt(), nextInt()
	a, b, c := make([]int, m), make([]int, m), make([]int, m)

	cost, use := make([][]int, n), make([][]bool, n)
	for i := 0; i < n; i++ {
		cost[i], use[i] = make([]int, n), make([]bool, n)
		for j := 0; j < n; j++ {
			cost[i][j] = inf
		}
	}
	for i := 0; i < m; i++ {
		a[i], b[i], c[i] = nextInt()-1, nextInt()-1, nextInt()
		cost[a[i]][b[i]], cost[b[i]][a[i]] = c[i], c[i]
	}

	// ワーシャルフロイド
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				cost[i][j] = min(cost[i][j], cost[i][k]+cost[k][j])
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		if c[i] > cost[a[i]][b[i]] {
			ans++
		}
	}

	puts(ans)
}
