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

type Sushi struct {
	x, v int
}

type Sushis []Sushi

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

	n, c := nextInt(), nextInt()
	ss := make(Sushis, n)
	for i := range ss {
		x, v := nextInt(), nextInt()
		ss[i] = Sushi{x, v}
	}

	// 時計回り，反時計回り
	cal1, cal2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		var dx int
		if i > 0 {
			cal1[i] = cal1[i-1]
			dx = ss[i].x - ss[i-1].x
		} else {
			dx = ss[i].x
		}
		cal1[i] += ss[i].v - dx
	}
	for i := n - 1; i >= 0; i-- {
		var dx int
		if i < n-1 {
			cal2[i] = cal2[i+1]
			dx = ss[i+1].x - ss[i].x
		} else {
			dx = c - ss[i].x
		}
		cal2[i] += ss[i].v - dx
	}

	mx1, mx2 := make([]int, n), make([]int, n)
	mx1[0], mx2[n-1] = 0, n-1
	for i := 1; i < n; i++ {
		mx1[i] = mx1[i-1]
		if cal1[i] > cal1[mx1[i]] {
			mx1[i] = i
		}
	}
	for i := n - 2; i >= 0; i-- {
		mx2[i] = mx2[i+1]
		if cal2[i] > cal2[mx2[i]] {
			mx2[i] = i
		}
	}

	// 途中で折り返す場合
	turn1, turn2 := 0, 0
	for i := 0; i < n-1; i++ {
		cal := cal1[i] + cal2[mx2[i+1]] - ss[i].x
		turn1 = max(turn1, cal)
	}
	for i := n - 1; i > 0; i-- {
		cal := cal2[i] + cal1[mx1[i-1]] - (c - ss[i].x)
		turn2 = max(turn2, cal)
	}

	puts(max(max(cal1...), max(cal2...), turn1, turn2))
}
