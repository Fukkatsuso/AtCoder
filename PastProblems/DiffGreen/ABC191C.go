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
	s := make([][]int, h)
	for i := range s {
		s[i] = make([]int, w)
		line := gets()
		for j := 0; j < w; j++ {
			if line[j] == '#' {
				s[i][j] = 1
			}
		}
	}

	ans := 0
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			black := s[i-1][j-1] + s[i-1][j] + s[i][j-1] + s[i][j]
			if black == 1 || black == 3 {
				ans++
			}
		}
	}

	puts(ans)
}
