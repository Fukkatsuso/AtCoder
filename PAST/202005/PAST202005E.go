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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, q := nextInt(), nextInt(), nextInt()
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		u, v := nextInt()-1, nextInt()-1
		edge[u][v], edge[v][u] = true, true
	}
	c := nextInts(n)

	for i := 0; i < q; i++ {
		s := nextInt()
		if s == 1 {
			x := nextInt() - 1
			puts(c[x])
			for j := range edge[x] {
				if edge[x][j] {
					c[j] = c[x]
				}
			}
		} else {
			x, y := nextInt()-1, nextInt()
			puts(c[x])
			c[x] = y
		}
	}
}
