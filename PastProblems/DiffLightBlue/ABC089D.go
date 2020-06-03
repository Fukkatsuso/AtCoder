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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, d := nextInt(), nextInt(), nextInt()
	// point[p]: pが書かれているマスの座標(y,x)
	point := make([][2]int, h*w+1)
	for i := 0; i < h; i++ {
		a := nextInts(w)
		for j := 0; j < w; j++ {
			point[a[j]] = [2]int{i, j}
		}
	}

	cost := make([]int, h*w+1)
	for i := d + 1; i <= h*w; i++ {
		cost[i] = cost[i-d] + abs(point[i][0]-point[i-d][0]) + abs(point[i][1]-point[i-d][1])
	}

	q := nextInt()
	for i := 0; i < q; i++ {
		l, r := nextInt(), nextInt()
		puts(cost[r] - cost[l])
	}
}
