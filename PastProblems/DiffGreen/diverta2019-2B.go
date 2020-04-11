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

func count(dist [][][]int, d []int) int {
	res := 0
	for i := range dist {
		for j := range dist[i] {
			if dist[i][j][0] == d[0] && dist[i][j][1] == d[1] {
				res++
			}
		}
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}

	dist := make([][][]int, n)
	for i := range dist {
		dist[i] = make([][]int, n)
		for j := range dist[i] {
			dist[i][j] = []int{x[j] - x[i], y[j] - y[i]}
		}
	}

	ans := n
	for i := range dist {
		for j := range dist[i] {
			if i == j {
				continue
			}
			d := []int{x[j] - x[i], y[j] - y[i]}
			ans = min(ans, n-count(dist, d))
		}
	}

	puts(ans)
}
