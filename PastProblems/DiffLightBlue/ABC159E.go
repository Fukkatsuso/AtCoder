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

	h, w, k := nextInt(), nextInt(), nextInt()
	s, sum := make([]string, h), make([][]int, h)
	for i := range s {
		s[i] = next()
		sum[i] = make([]int, w)
		for j := range sum[i] {
			sum[i][j] = int(s[i][j] - '0')
		}
	}
	// 累積和
	for i := 0; i < h; i++ {
		for j := 1; j < w; j++ {
			sum[i][j] += sum[i][j-1]
		}
	}
	for j := 0; j < w; j++ {
		for i := 1; i < h; i++ {
			sum[i][j] += sum[i-1][j]
		}
	}

	// ホワイトチョコ(1)の個数
	white := func(l, r, top, bottom int) int {
		res := sum[bottom][r]
		if l > 0 {
			res -= sum[bottom][l-1]
		}
		if top > 0 {
			res -= sum[top-1][r]
		}
		if l > 0 && top > 0 {
			res += sum[top-1][l-1]
		}
		return res
	}

	can := func(l, r int, btm []int) bool {
		top := 0
		for j := range btm {
			if white(l, r, top, btm[j]) > k {
				return false
			}
			top = btm[j] + 1
		}
		return true
	}

	ans := 10000
	// bitのibit目が1 => i:i+1で切る
	for bit := 0; bit < (1 << (h - 1)); bit++ {
		// 横に切ったときの各ブロックの底
		btm := make([]int, 0)
		for i, x := bit, 0; i > 0; i, x = i>>1, x+1 {
			if i&1 == 1 {
				btm = append(btm, x)
			}
		}
		btm = append(btm, h-1)
		ok, res := true, len(btm)-1
		// しゃくとり方で最小回数切る
		for l := 0; l < w; l++ {
			r := l - 1
			for r+1 < w && can(l, r+1, btm) {
				r++
			}
			if r < l {
				ok = false
				break
			}
			if r < w-1 {
				res++
			}
			l = r
		}
		if ok {
			ans = min(ans, res)
		}
	}

	puts(ans)
}
