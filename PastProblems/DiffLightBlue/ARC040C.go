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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	s := make([][]byte, n)
	for i := range s {
		s[i] = []byte(next())
	}

	ans := 0
	for i := 0; i < n; i++ {
		r := -1
		for j := 0; j < n; j++ {
			if s[i][j] == '.' {
				r = j
			}
		}
		if r < 0 {
			continue
		}
		ans++
		// 下の行を塗る
		if i == n-1 {
			break
		}
		for j := r; j < n; j++ {
			s[i+1][j] = 'o'
		}
	}

	puts(ans)
}
