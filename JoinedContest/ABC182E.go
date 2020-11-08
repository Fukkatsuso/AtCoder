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

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, n, m := nextInt4()
	cell := make([][]int, h)
	for i := range cell {
		cell[i] = make([]int, w)
	}
	for i := 0; i < n; i++ {
		a, b := nextInt()-1, nextInt()-1
		cell[a][b] = 1
	}
	for i := 0; i < m; i++ {
		c, d := nextInt()-1, nextInt()-1
		cell[c][d] = -1
	}

	// マス(i,j)の4方向に光源があるか
	l, r := make([][]bool, h), make([][]bool, h)
	u, d := make([][]bool, h), make([][]bool, h)
	for i := 0; i < h; i++ {
		l[i], r[i] = make([]bool, w), make([]bool, w)
		u[i], d[i] = make([]bool, w), make([]bool, w)
	}
	for i := 0; i < h; i++ {
		// l
		l[i][0] = (cell[i][0] > 0)
		for j := 1; j < w; j++ {
			l[i][j] = (cell[i][j] > 0) || (cell[i][j] == 0 && l[i][j-1])
		}
		// r
		r[i][w-1] = (cell[i][w-1] > 0)
		for j := w - 2; j >= 0; j-- {
			r[i][j] = (cell[i][j] > 0) || (cell[i][j] == 0 && r[i][j+1])
		}
	}
	for j := 0; j < w; j++ {
		// u
		u[0][j] = (cell[0][j] > 0)
		for i := 1; i < h; i++ {
			u[i][j] = (cell[i][j] > 0) || (cell[i][j] == 0 && u[i-1][j])
		}
		// d
		d[h-1][j] = (cell[h-1][j] > 0)
		for i := h - 2; i >= 0; i-- {
			d[i][j] = (cell[i][j] > 0) || (cell[i][j] == 0 && d[i+1][j])
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if l[i][j] || r[i][j] || u[i][j] || d[i][j] {
				ans++
			}
		}
	}
	puts(ans)
}
