package main

import (
	"bufio"
	"fmt"
	"math"
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

func min(nums ...float64) float64 {
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

func dist(x, y, px1, py1, px2, py2 int) float64 {
	dx, dy := px1-px2, py1-py2
	c := -dy*px1 + dx*py1
	d := float64(abs(x*dy-y*dx+c)) / math.Sqrt(float64(dx*dx+dy*dy))
	return d
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, y := nextInt(), nextInt()
	n := nextInt()
	px, py := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		px[i], py[i] = nextInt(), nextInt()
	}

	ans := 1000.0
	for i := 0; i < n; i++ {
		ans = min(ans, dist(x, y, px[i], py[i], px[(i+1)%n], py[(i+1)%n]))
	}

	puts(ans)
}
