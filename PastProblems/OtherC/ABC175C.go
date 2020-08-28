// 解説AC?
// d*kの乗算がオーバーフローしてWAになると気づいた

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

func nextInt() int64 {
	i, _ := strconv.Atoi(next())
	return int64(i)
}

func nextInt3() (int64, int64, int64) {
	return nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func min(nums ...int64) int64 {
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

	x, k, d := nextInt3()
	x = abs(x)

	var ans int64
	y := min(k, x/d)
	k -= y
	x -= y * d
	if k%2 == 0 {
		ans = x
	} else {
		ans = d - x
	}
	puts(ans)
}
