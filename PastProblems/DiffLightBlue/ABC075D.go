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

func min64(nums ...int64) int64 {
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

	N, K := nextInt(), nextInt()
	x, y := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}

	ans := int64(9223372036854775807)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				for l := 0; l < N; l++ {
					if i == j && j == k && k == l {
						continue
					}
					x1 := min(x[i], x[j], x[k], x[l])
					y1 := min(y[i], y[j], y[k], y[l])
					x2 := max(x[i], x[j], x[k], x[l])
					y2 := max(y[i], y[j], y[k], y[l])
					cnt := 0
					for p := 0; p < N; p++ {
						if x1 <= x[p] && x[p] <= x2 && y1 <= y[p] && y[p] <= y2 {
							cnt++
						}
					}
					if cnt >= K {
						s := int64(x2-x1) * int64(y2-y1)
						ans = min64(ans, s)
					}
				}
			}
		}
	}

	puts(ans)
}
