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

	n, k := nextInt(), nextInt()
	s := next()

	l, r := 0, 0
	zero := 0
	// 最初が0のときの例外処理
	if s[r] == '0' {
		zero = 1
	}
	ans := 0
	for l < n && r < n {
		// rを伸ばす
		for r < n && zero <= k {
			if r+1 < n && s[r] == '1' && s[r+1] == '0' {
				if zero == k {
					break
				}
				zero++
			}
			r++
		}
		if r == n {
			r--
		}
		// r-l+1人が並んでいる
		ans = max(ans, r-l+1)
		// 一番左の0区域を削除するまでlを進める
		for l < n && s[l] == '1' {
			l++
		}
		for l < n && s[l] == '0' {
			l++
		}
		zero--
	}

	puts(ans)
}
