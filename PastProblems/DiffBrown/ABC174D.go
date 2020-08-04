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

	n := nextInt()
	c := []byte(next())

	l, r := make([]int, n), make([]int, n)
	if c[0] == 'W' {
		l[0] = 1
	}
	for i := 1; i < n; i++ {
		if c[i] == 'W' {
			l[i] = l[i-1] + 1
		} else {
			l[i] = l[i-1]
		}
	}
	if c[n-1] == 'R' {
		r[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if c[i] == 'R' {
			r[i] = r[i+1] + 1
		} else {
			r[i] = r[i+1]
		}
	}

	ans := min(l[n-1], r[0])
	for i := 0; i < n-1; i++ {
		ans = min(ans, max(l[i], r[i+1]))
	}
	puts(ans)
}
