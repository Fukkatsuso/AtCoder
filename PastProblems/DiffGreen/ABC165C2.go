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

	n, m, q := nextInt(), nextInt(), nextInt()
	a, b, c, d := make([]int, q), make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		a[i], b[i], c[i], d[i] = nextInt()-1, nextInt()-1, nextInt(), nextInt()
	}

	point := func(A []int) int {
		p := 0
		for i := 0; i < q; i++ {
			if A[b[i]]-A[a[i]] == c[i] {
				p += d[i]
			}
		}
		return p
	}

	var dfs func(A []int, index int) int
	dfs = func(A []int, index int) int {
		if index == n {
			return point(A)
		}
		minI := 1
		if index > 0 {
			minI = A[index-1]
		}
		p := 0
		for i := minI; i <= m; i++ {
			A[index] = i
			p = max(p, dfs(A, index+1))
		}
		return p
	}

	puts(dfs(make([]int, n), 0))
}
