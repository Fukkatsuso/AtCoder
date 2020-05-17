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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	r, c, k := nextInt3()
	s := make([]string, r)
	for i := range s {
		s[i] = next()
	}

	top, below := make([][]int, r), make([][]int, r)
	for i := 0; i < r; i++ {
		top[i], below[i] = make([]int, c), make([]int, c)
	}
	for j := 0; j < c; j++ {
		for i := 0; i < r; i++ {
			if i > 0 {
				top[i][j] = top[i-1][j]
				below[r-1-i][j] = below[r-i][j]
			}
			if s[i][j] == 'o' {
				top[i][j]++
			} else {
				top[i][j] = 0
			}
			if s[r-1-i][j] == 'o' {
				below[r-1-i][j]++
			} else {
				below[r-1-i][j] = 0
			}
		}
	}

	ans := 0
	for i := k - 1; i <= r-k; i++ {
		for j := k - 1; j <= c-k; j++ {
			ok := true
			// left
			for l := 0; ok && l < k; l++ {
				if top[i][j-l] < k-l || below[i][j-l] < k-l {
					ok = false
				}
			}
			// right
			for l := 1; ok && l < k; l++ {
				if top[i][j+l] < k-l || below[i][j+l] < k-l {
					ok = false
				}
			}
			if ok {
				ans++
			}
		}
	}

	puts(ans)
}
