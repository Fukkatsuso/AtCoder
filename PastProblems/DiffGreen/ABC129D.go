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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
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

	// 前処理
	yoko, tate := make([][]int, h), make([][]int, h)
	for i := 0; i < h; i++ {
		yoko[i], tate[i] = make([]int, w), make([]int, w)
	}
	// yoko
	for i := 0; i < h; i++ {
		for l := 0; l < w; {
			for l < w && s[i][l] == '#' {
				l++
			}
			r := l
			for r+1 < w && s[i][r+1] == '.' {
				r++
			}
			dist := r - l + 1
			for j := l; j <= r && j < w; j++ {
				yoko[i][j] = dist
			}
			l = r + 1
		}
	}
	// tate
	for i := 0; i < w; i++ {
		for u := 0; u < h; {
			for u < h && s[u][i] == '#' {
				u++
			}
			d := u
			for d+1 < h && s[d+1][i] == '.' {
				d++
			}
			dist := d - u + 1
			for j := u; j <= d && j < h; j++ {
				tate[j][i] = dist
			}
			u = d + 1
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				ans = max(ans, yoko[i][j]+tate[i][j]-1)
			}
		}
	}

	puts(ans)
}
