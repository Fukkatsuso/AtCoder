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

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
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

func divFloor(a, b int) int {
	return (a + (b - 1)) / b
}

func cost(n, a, b, c, d int) int {
	dp := map[int]int{}
	dp[0], dp[1] = 0, d
	var rec func(x int) int
	rec = func(x int) int {
		if v, ok := dp[x]; ok {
			return v
		}
		res := 1 << 60
		if x < res/d {
			res = d * x
		}
		res = min(
			res,
			d*(x-(x/2)*2)+a+rec(x/2),
			d*(divFloor(x, 2)*2-x)+a+rec(divFloor(x, 2)),
			d*(x-(x/3)*3)+b+rec(x/3),
			d*(divFloor(x, 3)*3-x)+b+rec(divFloor(x, 3)),
			d*(x-(x/5)*5)+c+rec(x/5),
			d*(divFloor(x, 5)*5-x)+c+rec(divFloor(x, 5)),
		)
		dp[x] = res
		return res
	}
	return rec(n)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	t := nextInt()

	for i := 0; i < t; i++ {
		n := nextInt()
		a, b, c, d := nextInt4()
		ans := cost(n, a, b, c, d)
		puts(ans)
	}
}
