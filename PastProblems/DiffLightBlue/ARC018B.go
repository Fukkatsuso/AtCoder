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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	ps := make([]Point, n)
	for i := range ps {
		x, y := nextInt(), nextInt()
		ps[i] = Point{x, y}
	}

	ok := func(pa, pb, pc Point) bool {
		ss := (pa.x-pc.x)*(pb.y-pc.y) - (pb.x-pc.x)*(pa.y-pc.y)
		return ss%2 == 0
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				// 一直線上にある場合
				if (ps[i].x-ps[k].x)*(ps[j].y-ps[k].y) == (ps[j].x-ps[k].x)*(ps[i].y-ps[k].y) {
					continue
				}
				if ok(ps[i], ps[j], ps[k]) {
					ans++
				}
			}
		}
	}
	puts(ans)
}
