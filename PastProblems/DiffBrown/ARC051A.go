package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, y := make([]int, 3), make([]int, 3)
	x[0], y[0] = nextInt(), nextInt()
	r := nextInt()
	x[1], y[1], x[2], y[2] = nextInt(), nextInt(), nextInt(), nextInt()

	red, blue := false, false

	// redがblueを覆う
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			blue = blue || float64(r) < dist(x[0], y[0], x[i], y[j])
		}
	}

	// blueがredを覆う
	p := [][]int{
		[]int{x[0] + r, y[0] + r},
		[]int{x[0] + r, y[0] - r},
		[]int{x[0] - r, y[0] + r},
		[]int{x[0] - r, y[0] - r},
	}
	for _, q := range p {
		red = red || !(x[1] <= q[0] && q[0] <= x[2] && y[1] <= q[1] && q[1] <= y[2])
	}

	YesOrNo(red)
	YesOrNo(blue)
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

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func dist(x1, y1, x2, y2 int) float64 {
	dx := float64(x1 - x2)
	dy := float64(y1 - y2)
	return math.Sqrt(dx*dx + dy*dy)
}

func YesOrNo(ok bool) {
	if ok {
		puts("YES")
	} else {
		puts("NO")
	}
}
