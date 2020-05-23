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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
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

	n, k := nextInt2()
	s := nextInts(n)
	for i := range s {
		if s[i] == 0 {
			puts(n)
			return
		}
	}

	ans := 0
	pow := 1
	for l, r := 0, 0; l < n; l++ {
		// rを限界まで伸ばす
		for r < n && pow*s[r] <= k {
			pow *= s[r]
			r++
		}
		ans = max(ans, r-l)
		if r == l {
			r++
		} else {
			pow /= s[l]
		}
	}

	puts(ans)
}
