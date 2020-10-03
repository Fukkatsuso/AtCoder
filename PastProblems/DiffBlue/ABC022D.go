// 解説AC

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

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func dist(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}

func max(nums ...float64) float64 {
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

	ax, ay := make([]float64, n), make([]float64, n)
	// 1日目の星の重心
	sax, say := 0.0, 0.0
	for i := 0; i < n; i++ {
		ax[i], ay[i] = nextFloat64(), nextFloat64()
		sax += ax[i]
		say += ay[i]
	}
	sax /= float64(n)
	say /= float64(n)

	bx, by := make([]float64, n), make([]float64, n)
	// 2日目の星の重心
	sbx, sby := 0.0, 0.0
	for i := 0; i < n; i++ {
		bx[i], by[i] = nextFloat64(), nextFloat64()
		sbx += bx[i]
		sby += by[i]
	}
	sbx /= float64(n)
	sby /= float64(n)

	// 重心からの最遠距離
	da, db := 0.0, 0.0
	for i := 0; i < n; i++ {
		da = max(da, dist(sax, say, ax[i], ay[i]))
		db = max(db, dist(sbx, sby, bx[i], by[i]))
	}
	p := db / da
	puts(p)
}
