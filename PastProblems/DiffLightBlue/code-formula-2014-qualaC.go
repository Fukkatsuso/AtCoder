// 解説AC
// 問題文の意味がつかめなかった
// 参照 https://kmjp.hatenablog.jp/entry/2014/08/21/0930

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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	a := make([][]int, n)
	for i := range a {
		a[i] = nextInts(k)
	}

	nk := n * k
	rank := make([]int, nk+1)
	invited := map[int]bool{}
	rem := k
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			rank[j*n+i] = a[i][j]
		}

		x := 0
		win := make([]int, 0)
		for y := 0; y < nk; y++ {
			if invited[rank[y]] {
				continue
			}
			if x++; x > rem {
				break
			}
			if rank[y] > 0 {
				win = append(win, rank[y])
				invited[rank[y]] = true
			}
		}

		sort.Sort(sort.IntSlice(win))
		for j := 0; j < len(win); j++ {
			if j == len(win)-1 {
				putf("%d", win[j])
			} else {
				putf("%d ", win[j])
			}
		}
		puts()
		rem -= len(win)
	}
}

// テストケース

// 2 11
// 1 2 3 4 5 6 7 8 9 10 11
// 1 2 15 14 13 16 17 18 19 20 21

// 4 5
// 1 2 3 4 5
// 2 1 3 4 5
// 1 2 3 4 5
// 2 1 3 4 5
