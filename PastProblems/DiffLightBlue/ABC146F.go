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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
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

	n, m := nextInt(), nextInt()
	s := next()

	// can[i]: マスi以降にある，進めるマス
	can := make([]int, n+1)
	can[n] = n
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			can[i] = can[i+1]
		} else {
			can[i] = i
		}
	}

	ans := make([]int, 0)
	from, to := n, n
	for from > 0 {
		to = can[max(0, from-m)]
		if to == from {
			puts(-1)
			return
		}
		ans = append(ans, from-to)
		from = to
	}

	for i := len(ans) - 1; i > 0; i-- {
		putf("%d ", ans[i])
	}
	puts(ans[0])
}
