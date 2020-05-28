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

	n := nextInt()
	a := nextInts(n)

	ans1 := 0
	for i, sum, pos := 0, 0, true; i < n; i, pos = i+1, !pos {
		var d int
		if pos && (sum+a[i]) <= 0 {
			d = abs(1 - (sum + a[i]))
			ans1 += d
		} else if !pos && (sum+a[i]) >= 0 {
			d = -abs(-1 - (sum + a[i]))
			ans1 -= d
		}
		sum += a[i] + d
	}

	ans2 := 0
	for i, sum, pos := 0, 0, false; i < n; i, pos = i+1, !pos {
		var d int
		if pos && (sum+a[i]) <= 0 {
			d = abs(1 - (sum + a[i]))
			ans2 += d
		} else if !pos && (sum+a[i]) >= 0 {
			d = -abs(-1 - (sum + a[i]))
			ans2 -= d
		}
		sum += a[i] + d
	}

	puts(min(ans1, ans2))
}
