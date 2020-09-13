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

type Point struct {
	x, y int
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
	pp := make([]Point, n)
	for i := range pp {
		x, y := nextInt(), nextInt()
		pp[i] = Point{x, y}
	}

	add, sub := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		add[i] = pp[i].x + pp[i].y
		sub[i] = pp[i].x - pp[i].y
	}

	ans := max(
		max(add...)-min(add...),
		-min(sub...)+max(sub...),
		max(sub...)-max(sub...),
		-min(add...)+max(add...),
	)

	puts(ans)
}
