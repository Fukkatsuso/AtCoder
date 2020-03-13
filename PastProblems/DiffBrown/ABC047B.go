package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	w, h, n := nextInt(), nextInt(), nextInt()

	up, down, right, left := h, 0, w, 0
	for i := 0; i < n; i++ {
		x, y, a := nextInt(), nextInt(), nextInt()
		switch a {
		case 1:
			left = max(left, x)
		case 2:
			right = min(right, x)
		case 3:
			down = max(down, y)
		case 4:
			up = min(up, y)
		}
	}

	ans := 0
	if left <= right && down <= up {
		ans = (right - left) * (up - down)
	}
	puts(ans)
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
