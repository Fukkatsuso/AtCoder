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

func firstWin(win, visited [][]bool, s []string, h, w, i, j int) bool {
	if visited[i][j] {
		return win[i][j]
	}
	visited[i][j] = true
	di := []int{1, 1, 0}
	dj := []int{0, 1, 1}
	ok := false
	for k := 0; k < 3; k++ {
		nextI, nextJ := i+di[k], j+dj[k]
		if nextI < h && nextJ < w && s[nextI][nextJ] == '.' {
			ok = ok || !firstWin(win, visited, s, h, w, nextI, nextJ)
		}
	}
	win[i][j] = ok
	return win[i][j]
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

	win, visited := make([][]bool, h), make([][]bool, h)
	for i := range win {
		win[i], visited[i] = make([]bool, w), make([]bool, w)
	}

	if firstWin(win, visited, s, h, w, 0, 0) {
		puts("First")
	} else {
		puts("Second")
	}
}
