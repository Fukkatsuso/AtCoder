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

func dfs(group []int, cross [][]int, p, num int) {
	var rec func(p int)
	rec = func(p int) {
		for _, to := range cross[p] {
			if group[to] == 0 {
				group[to] = num
				rec(to)
			}
		}
	}
	rec(p)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	cross := make([][]int, n)
	for i := range cross {
		cross[i] = make([]int, 0)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		cross[a] = append(cross[a], b)
		cross[b] = append(cross[b], a)
	}

	group := make([]int, n)
	num := 0
	for i := 0; i < n; i++ {
		if group[i] == 0 {
			num++
			dfs(group, cross, i, num)
		}
	}

	puts(num - 1)
}
