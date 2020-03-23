package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isff(f [][]bool, a, b int) bool {
	if f[a][b] || a == b {
		return false
	}
	for i, n := 0, len(f); i < n; i++ {
		if f[a][i] && f[i][b] {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		f[a][b] = true
		f[b][a] = true
	}

	// 友達の友達か
	ff := make([][]bool, n)
	for i := range ff {
		ff[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ff[i][j] = isff(f, i, j)
		}
	}

	for i := range ff {
		cnt := 0
		for j := range ff[i] {
			if ff[i][j] {
				cnt++
			}
		}
		puts(cnt)
	}
}

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
