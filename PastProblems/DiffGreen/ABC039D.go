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

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := range s {
		s[i] = next()
	}

	// true: #, false: .
	t := make([][]bool, h)
	for i := range t {
		t[i] = make([]bool, w)
	}

	dx := []int{0, 1, 0, -1, 1, 1, -1, -1}
	dy := []int{-1, 0, 1, 0, -1, 1, 1, -1}

	// 8方を囲まれた#は確定
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				continue
			}
			ok := true
			for k := 0; k < 8; k++ {
				x, y := j+dx[k], i+dy[k]
				if x < 0 || w <= x || y < 0 || h <= y {
					continue
				}
				ok = ok && (s[y][x] == '#')
			}
			t[i][j] = ok
		}
	}

	can := true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' || t[i][j] {
				continue
			}
			ok := false
			// #が確定していないマス
			for k := 0; k < 8; k++ {
				x, y := j+dx[k], i+dy[k]
				if x < 0 || w <= x || y < 0 || h <= y {
					continue
				}
				ok = ok || t[y][x]
			}
			can = can && ok
		}
	}

	if !can {
		puts("impossible")
		return
	}

	puts("possible")
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if t[i][j] {
				putf("#")
			} else {
				putf(".")
			}
		}
		puts()
	}
}
