package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, x, y := nextInt(), nextInt()-1, nextInt()-1

	ans := map[int]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := min(j-i, abs(x-i)+1+abs(y-j))
			ans[d]++
		}
	}

	for i := 1; i < n; i++ {
		puts(ans[i])
	}
}

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
