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

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	n, a, b, c := nextInt4()
	l := nextInts(n)

	var dfs func(x, y, z, i, mp int) int
	dfs = func(x, y, z, i, mp int) int {
		if i == n {
			if x*y*z == 0 {
				return 1 << 50
			}
			return mp + abs(a-x) + abs(b-y) + abs(c-z)
		}

		res := dfs(x, y, z, i+1, mp)
		if x == 0 {
			res = min(res, dfs(x+l[i], y, z, i+1, mp))
		} else {
			res = min(res, dfs(x+l[i], y, z, i+1, mp+10))
		}
		if y == 0 {
			res = min(res, dfs(x, y+l[i], z, i+1, mp))
		} else {
			res = min(res, dfs(x, y+l[i], z, i+1, mp+10))
		}
		if z == 0 {
			res = min(res, dfs(x, y, z+l[i], i+1, mp))
		} else {
			res = min(res, dfs(x, y, z+l[i], i+1, mp+10))
		}
		return res
	}

	puts(dfs(0, 0, 0, 0, 0))
}
