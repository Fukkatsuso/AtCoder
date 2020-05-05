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

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
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

// ピラミッドの高さ
func pyramid(cx, cy, x, y, h int) int {
	return h + abs(x-cx) + abs(y-cy)
}

// 地点(x, y)の高さ
func height(cx, cy, h, x, y int) int {
	return max(0, h-abs(x-cx)-abs(y-cy))
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

	n := nextInt()
	x, y, h := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i], h[i] = nextInt3()
	}

	for cx := 0; cx <= 100; cx++ {
		for cy := 0; cy <= 100; cy++ {
			ok := true
			k := 0
			for k < n && h[k] == 0 {
				k++
			}
			H := max(1, pyramid(cx, cy, x[k%n], y[k%n], h[k%n]))
			for l := 0; ok && l < n; l++ {
				ok = ok && (max(height(cx, cy, H, x[l], y[l]), 0) == h[l])
			}

			if ok {
				puts(cx, cy, H)
				return
			}
		}
	}
}
