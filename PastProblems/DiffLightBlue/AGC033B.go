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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, n := nextInt(), nextInt(), nextInt()
	sr, sc := nextInt()-1, nextInt()-1
	s, t := next(), next()

	// 後手が勝てる駒の範囲
	// タテ方向
	u, d := 0, h
	if s[n-1] == 'U' {
		u++
	} else if s[n-1] == 'D' {
		d--
	}
	for i := n - 2; i >= 0; i-- {
		if t[i] == 'U' {
			d = min(d+1, h)
		} else if t[i] == 'D' {
			u = max(0, u-1)
		}
		if s[i] == 'U' {
			u++
		} else if s[i] == 'D' {
			d--
		}
		if u >= d {
			puts("NO")
			return
		}
	}
	if sr < u || d <= sr {
		puts("NO")
		return
	}

	// ヨコ方向
	l, r := 0, w
	if s[n-1] == 'L' {
		l++
	} else if s[n-1] == 'R' {
		r--
	}
	for i := n - 2; i >= 0; i-- {
		if t[i] == 'L' {
			r = min(r+1, w)
		} else if t[i] == 'R' {
			l = max(0, l-1)
		}
		if s[i] == 'L' {
			l++
		} else if s[i] == 'R' {
			r--
		}
		if l >= r {
			puts("NO")
			return
		}
	}
	if sc < l || r <= sc {
		puts("NO")
		return
	}

	puts("YES")
}
