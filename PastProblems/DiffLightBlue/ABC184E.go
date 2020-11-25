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

	h, w := getInt(), getInt()
	a := make([]string, h)
	var si, sj, gi, gj int
	// テレポーターのマスのリスト
	tp := make([][][2]int, 26)
	for i := range tp {
		tp[i] = make([][2]int, 0)
	}
	for i := range a {
		a[i] = gets()
		for j, c := range a[i] {
			if c == 'S' {
				si, sj = i, j
			} else if c == 'G' {
				gi, gj = i, j
			} else if c != '.' && c != '#' {
				tp[int(c-'a')] = append(tp[int(c-'a')], [2]int{i, j})
			}
		}
	}

	t := make([][]int, h)
	for i := range t {
		t[i] = make([]int, w)
		for j := range t[i] {
			t[i][j] = -1
		}
	}
	t[si][sj] = 0

	q := make([][2]int, 0)
	q = append(q, [2]int{si, sj})
	used := make([]bool, 26)
	di := []int{1, 0, -1, 0}
	dj := []int{0, 1, 0, -1}
	for len(q) > 0 {
		i, j := q[0][0], q[0][1]
		q = q[1:]
		// テレポート
		if c := a[i][j]; 'a' <= c && c <= 'z' && !used[int(c-'a')] {
			for _, tpi := range tp[int(c-'a')] {
				// テレポート先が未訪問
				if t[tpi[0]][tpi[1]] < 0 {
					t[tpi[0]][tpi[1]] = t[i][j] + 1
					q = append(q, tpi)
				}
			}
			used[int(c-'a')] = true
		}
		// 隣り合うマスへ移動
		for k := 0; k < 4; k++ {
			nextI, nextJ := i+di[k], j+dj[k]
			if nextI < 0 || h <= nextI || nextJ < 0 || w <= nextJ {
				continue
			}
			if a[nextI][nextJ] == '#' || t[nextI][nextJ] >= 0 {
				continue
			}
			t[nextI][nextJ] = t[i][j] + 1
			q = append(q, [2]int{nextI, nextJ})
		}
	}

	puts(t[gi][gj])
}
