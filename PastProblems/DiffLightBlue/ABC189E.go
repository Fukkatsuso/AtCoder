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

func MatixProduct(m1, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i := range res {
		res[i] = make([]int, len(m2[i]))
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			for k := range m1[i] {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	// 反時計回り
	rot := func() [][]int {
		return [][]int{
			{0, -1, 0},
			{1, 0, 0},
			{0, 0, 1},
		}
	}
	// 時計回り
	rotRev := func() [][]int {
		return [][]int{
			{0, 1, 0},
			{-1, 0, 0},
			{0, 0, 1},
		}
	}
	// x=pについて線対称
	symmX := func(p int) [][]int {
		return [][]int{
			{-1, 0, 2 * p},
			{0, 1, 0},
			{0, 0, 1},
		}
	}
	// y=pについて線対称
	symmY := func(p int) [][]int {
		return [][]int{
			{1, 0, 0},
			{0, -1, 2 * p},
			{0, 0, 1},
		}
	}

	n := getInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = getInt(), getInt()
	}

	mats := make([][][]int, 0)
	mat := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	mats = append(mats, mat)
	m := getInt()
	for i := 0; i < m; i++ {
		op := getInt()
		if op == 1 {
			mat = MatixProduct(rotRev(), mat)
		} else if op == 2 {
			mat = MatixProduct(rot(), mat)
		} else {
			p := getInt()
			if op == 3 {
				mat = MatixProduct(symmX(p), mat)
			} else {
				mat = MatixProduct(symmY(p), mat)
			}
		}
		mats = append(mats, mat)
	}

	q := getInt()
	for i := 0; i < q; i++ {
		a, b := getInt(), getInt()-1
		xy := MatixProduct(mats[a], [][]int{{x[b]}, {y[b]}, {1}})
		puts(xy[0][0], xy[1][0])
	}
}
