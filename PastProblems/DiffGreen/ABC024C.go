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

	_, d, k := nextInt(), nextInt(), nextInt()
	l, r := make([]int, d), make([]int, d)
	for i := 0; i < d; i++ {
		l[i], r[i] = nextInt(), nextInt()
	}
	s, t := make([]int, k), make([]int, k)
	for i := 0; i < k; i++ {
		s[i], t[i] = nextInt(), nextInt()
	}
	
	for i := 0; i < k; i++ {
		// 日数
		ans := 0
		for ; ans < d && s[i] != t[i]; ans++ {
			if !(l[ans]<=s[i] && s[i]<=r[ans]) {
				continue
			}
			if s[i] < t[i] {
				s[i] = min(t[i], r[ans])
			} else if s[i] > t[i] {
				s[i] = max(t[i], l[ans])
			}
		}
		puts(ans)
	}
}
