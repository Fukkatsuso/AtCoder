package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
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

func nextPermutation(a []int) bool {
	n := len(a)
	reverse := func(begin, end int) {
		for i := 0; i < (end-begin)/2; i++ {
			a[begin+i], a[end-1-i] = a[end-1-i], a[begin+i]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			var j int
			for j = n - 1; a[i] >= a[j]; j-- {
			}
			a[i], a[j] = a[j], a[i]
			reverse(i+1, n)
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, M, R := nextInt(), nextInt(), nextInt()
	r := make([]int, R)
	for i := range r {
		r[i] = nextInt() - 1
	}
	sort.Sort(sort.IntSlice(r))
	cost := make([][]int, N)
	for i := range cost {
		cost[i] = make([]int, N)
		for j := range cost[i] {
			cost[i][j] = 1 << 60
		}
	}
	for i := 0; i < M; i++ {
		a, b, c := nextInt()-1, nextInt()-1, nextInt()
		cost[a][b], cost[b][a] = c, c
	}

	// ワーシャルフロイド
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				cost[i][j] = min(cost[i][j], cost[i][k]+cost[k][j])
			}
		}
	}

	var ans int
	for i := 0; i < R-1; i++ {
		ans += cost[r[i]][r[i+1]]
	}
	// 順列全探索
	for nextPermutation(r) {
		x := 0
		for i := 0; i < R-1; i++ {
			x += cost[r[i]][r[i+1]]
		}
		ans = min(ans, x)
	}

	puts(ans)
}
