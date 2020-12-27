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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	l := getInt()

	r := 0
	for (1<<(r+1)) <= l && r+1 <= 19 {
		r++
	}

	n, m := r+1, 0
	u, v, w := make([]int, 0), make([]int, 0), make([]int, 0)

	add := func(from, to, length int) {
		u = append(u, from)
		v = append(v, to)
		w = append(w, length)
		m++
	}

	// 頂点iから頂点i+1へ，長さ0の辺と長さ1<<iの辺を張る
	for i := 1; i < n; i++ {
		add(i, i+1, 0)
		add(i, i+1, 1<<(i-1))
	}
	// 足りない分を調整
	for i := n - 1; i >= 1; i-- {
		if l-(1<<(i-1)) >= (1 << r) {
			add(i, n, l-(1<<(i-1)))
			l -= 1 << (i - 1)
		}
	}

	puts(n, m)
	for i := 0; i < m; i++ {
		puts(u[i], v[i], w[i])
	}
}
