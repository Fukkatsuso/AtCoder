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
	c := make([][]int, h)
	// ホワイトとブラックを異符号にする
	for i := range c {
		c[i] = make([]int, w)
		for j := range c[i] {
			c[i][j] = nextInt()
			if (i+j)%2 == 1 {
				c[i][j] *= -1
			}
		}
	}

	// 2次元累積和
	d := make([][]int, h)
	for i := 0; i < h; i++ {
		d[i] = make([]int, w)
		for j := 0; j < w; j++ {
			if j > 0 {
				d[i][j] = d[i][j-1]
			}
			d[i][j] += c[i][j]
		}
	}
	for j := 0; j < w; j++ {
		for i := 1; i < h; i++ {
			d[i][j] += d[i-1][j]
		}
	}

	// 長方形領域の濃度の合計
	calc := func(t, b, l, r int) int {
		res := d[b][r]
		if t > 0 {
			res -= d[t-1][r]
		}
		if l > 0 {
			res -= d[b][l-1]
		}
		if t > 0 && l > 0 {
			res += d[t-1][l-1]
		}
		return res
	}

	ans := 0
	for t := 0; t < h; t++ {
		for b := t; b < h; b++ {
			for l := 0; l < w; l++ {
				for r := l; r < w; r++ {
					if calc(t, b, l, r) != 0 {
						continue
					}
					ans = max(ans, (b-t+1)*(r-l+1))
				}
			}
		}
	}

	puts(ans)
}
